package valdo

import (
	"github.com/orsinium-labs/jsony"
)

type ObjectType struct {
	ps       []PropertyType
	cs       []Constraint[map[string]any]
	extra    bool
	extraVal Validator
}

var _ Validator = Object()

func Object(ps ...PropertyType) ObjectType {
	return ObjectType{ps: ps}
}

func (obj ObjectType) Constrain(cs ...Constraint[map[string]any]) ObjectType {
	obj.cs = append(obj.cs, cs...)
	return obj
}

func (obj ObjectType) AllowExtra(v Validator) ObjectType {
	obj.extra = true
	obj.extraVal = v
	return obj
}

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
	for _, p := range obj.ps {
		val, found := data[p.name]
		if !found {
			if !p.optional {
				res.Add(ErrRequired{Name: p.name})
			}
			continue
		}
		res.Add(p.Validate(val))
	}
	for _, c := range obj.cs {
		res.Add(c.check(data))
	}
	if obj.extraVal != nil {
		for name, val := range data {
			if !obj.hasProperty(name) {
				err := obj.extraVal.Validate(val)
				res.Add(ErrProperty{Name: name, Err: err})
			}
		}
	}
	if !obj.extra {
		for name := range data {
			if !obj.hasProperty(name) {
				res.Add(ErrUnexpected{Name: name})
			}
		}
	}
	return res.Flatten()
}

func (obj ObjectType) hasProperty(name string) bool {
	for _, p := range obj.ps {
		if p.name == name {
			return true
		}
	}
	return false
}

func (obj ObjectType) Schema() jsony.Object {
	required := make(jsony.Array[jsony.String], 0)
	properties := make(jsony.Map, 0)
	for _, p := range obj.ps {
		if !p.optional {
			required = append(required, jsony.String(p.name))
		}
		properties[jsony.String(p.name)] = p.validator.Schema()
	}
	res := jsony.Object{
		jsony.Field{K: "type", V: jsony.SafeString("object")},
	}
	if len(properties) > 0 {
		res = append(res, jsony.Field{K: "properties", V: properties})
	}
	if len(required) != 0 {
		res = append(res, jsony.Field{K: "required", V: required})
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

type PropertyType struct {
	name      string
	validator Validator
	optional  bool
}

func Property(name string, v Validator) PropertyType {
	return PropertyType{name: name, validator: v}
}

func (p PropertyType) Optional() PropertyType {
	p.optional = true
	return p
}

func (p PropertyType) Validate(data any) Error {
	err := p.validator.Validate(data)
	if err != nil {
		return ErrProperty{Name: p.name, Err: err}
	}
	return nil
}
