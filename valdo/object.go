package valdo

import (
	"regexp"

	"github.com/orsinium-labs/jsony"
)

// ObjectType is constructed by [Object].
type ObjectType struct {
	ps       []PropertyType
	cs       []Constraint[map[string]any]
	extra    bool
	extraVal Validator
}

func Map(value Validator, cs ...Constraint[map[string]any]) ObjectType {
	return Object().AllowExtra(value).Constrain(cs...)
}

// Object maps to "object" in JSON and "struct" or "map" in Go.
func Object(ps ...PropertyType) ObjectType {
	return ObjectType{ps: ps}
}

// Add constraints to the object.
//
// One of:
//
//   - [PropertyNames]
//   - [MinProperties]
//   - [MaxProperties]
func (obj ObjectType) Constrain(cs ...Constraint[map[string]any]) ObjectType {
	obj.cs = append(obj.cs, cs...)
	return obj
}

// Allow additional properties.
//
// The validator, if not nil, will be used to validate the additional properties.
//
// https://json-schema.org/understanding-json-schema/reference/object#additionalproperties
func (obj ObjectType) AllowExtra(v Validator) ObjectType {
	obj.extra = true
	obj.extraVal = v
	return obj
}

// Validate implements [Validator].
func (obj ObjectType) Validate(data any) Error {
	switch d := data.(type) {
	case map[string]any:
		return obj.validateMap(d)
	default:
		return ErrType{Got: getTypeName(data), Expected: "object"}
	}
}

func (obj ObjectType) validateMap(data map[string]any) Error {
	if data == nil {
		return ErrType{Got: "null", Expected: "object"}
	}
	res := Errors{}
	handledNames := map[string]struct{}{}
	for _, p := range obj.ps {
		if p.rex != nil {
			for name, val := range data {
				if !p.rex.MatchString(name) {
					continue
				}
				handledNames[name] = struct{}{}
				res.Add(p.validate(val))
			}
			continue
		}

		val, found := data[p.name]
		if !found {
			if !p.optional {
				res.Add(ErrRequired{Name: p.name})
			}
			continue
		}
		handledNames[p.name] = struct{}{}
		res.Add(p.validate(val))
		if len(p.depReq) > 0 {
			for _, name := range p.depReq {
				_, found := data[name]
				if !found {
					res.Add(ErrRequired{Name: name})
				}
			}
		}
	}
	for _, c := range obj.cs {
		res.Add(c.check(data))
	}
	if obj.extraVal != nil {
		for name, val := range data {
			_, handled := handledNames[name]
			if !handled {
				err := obj.extraVal.Validate(val)
				if err != nil {
					res.Add(ErrProperty{Name: name, Err: err})
				}
			}
		}
	} else if !obj.extra {
		for name := range data {
			_, handled := handledNames[name]
			if !handled {
				res.Add(ErrUnexpected{Name: name})
			}
		}
	}
	return res.Flatten()
}

// Schema implements [Validator].
func (obj ObjectType) Schema() jsony.Object {
	required := make(jsony.Array[jsony.String], 0)
	properties := make(jsony.Map, 0)
	patternProps := make(jsony.Map, 0)
	for _, p := range obj.ps {
		name := jsony.String(p.name)
		if !p.optional {
			required = append(required, name)
		}
		pSchema := p.validator.Schema()
		if len(pSchema) > 0 {
			if p.rex != nil {
				patternProps[name] = pSchema
			} else {
				properties[name] = pSchema
			}
		}
	}
	res := jsony.Object{
		jsony.Field{K: "type", V: jsony.SafeString("object")},
	}
	if len(properties) > 0 {
		res = append(res, jsony.Field{K: "properties", V: properties})
	}
	if len(patternProps) > 0 {
		res = append(res, jsony.Field{K: "patternProperties", V: patternProps})
	}
	if len(required) != 0 {
		res = append(res, jsony.Field{K: "required", V: required})
	}

	depReq := make(jsony.Map)
	for _, p := range obj.ps {
		if len(p.depReq) > 0 {
			val := make(jsony.Array[jsony.String], len(p.depReq))
			for i, reqName := range p.depReq {
				val[i] = jsony.String(reqName)
			}
			depReq[jsony.String(p.name)] = val
		}
	}
	if len(depReq) > 0 {
		res = append(res, jsony.Field{K: "dependentRequired", V: depReq})
	}

	if obj.extraVal != nil {
		res = append(res, jsony.Field{K: "additionalProperties", V: obj.extraVal.Schema()})
	} else if !obj.extra {
		res = append(res, jsony.Field{K: "additionalProperties", V: jsony.False})
	}
	for _, c := range obj.cs {
		res = append(res, c.field)
	}
	return res
}

// PropertyType is constructed by [Property].
type PropertyType struct {
	name      string
	rex       *regexp.Regexp
	validator Validator
	optional  bool
	depReq    []string
}

// Property is a key-value pair of an [Object].
func Property(name string, v Validator) PropertyType {
	var rex *regexp.Regexp
	if name != "" && name[0] == '^' {
		rex = regexp.MustCompile(name)
	}
	return PropertyType{name: name, validator: v, rex: rex}
}

// Mark the property as optional.
//
// By default, all properties listed in the object are required.
//
// https://json-schema.org/understanding-json-schema/reference/object#required
func (p PropertyType) Optional() PropertyType {
	p.optional = true
	return p
}

// If the property is present, require also the given properties to be present.
//
// https://json-schema.org/understanding-json-schema/reference/conditionals#dependentRequired
func (p PropertyType) AlsoRequire(name string, names ...string) PropertyType {
	if p.rex != nil {
		panic("pattern properties cannot be required")
	}
	p.depReq = append(p.depReq, name)
	p.depReq = append(p.depReq, names...)
	return p
}

func (p PropertyType) validate(data any) Error {
	err := p.validator.Validate(data)
	if err != nil {
		return ErrProperty{Name: p.name, Err: err}
	}
	return nil
}
