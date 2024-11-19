package valdo

import (
	"reflect"

	"github.com/orsinium-labs/jsony"
)

type Object struct {
	ps    []Property
	extra bool
}

var _ Validator = O()

func O(ps ...Property) Object {
	return Object{ps: ps}
}

func (obj Object) Validate(data any) Error {
	switch d := data.(type) {
	case map[string]any:
		return obj.validateMap(d)
	default:
		return ErrType{Got: getTypeName(data), Expected: "object"}
	}
}

func (obj Object) validateMap(data map[string]any) Error {
	res := Errors{}
	for _, p := range obj.ps {
		val, found := data[p.name]
		if !found {
			if p.optional {
				continue
			}
			res.Errs = append(res.Errs, ErrRequired{Name: p.name})
		}
		err := p.Validate(val)
		if err != nil {
			res.Errs = append(res.Errs, err)
		}
	}
	if !obj.extra {
		for name := range data {
			if !obj.hasProperty(name) {
				res.Errs = append(res.Errs, ErrUnexpected{Name: name})
			}
		}
	}
	return res
}

func (obj Object) hasProperty(name string) bool {
	for _, p := range obj.ps {
		if p.name == name {
			return true
		}
	}
	return false
}

// func (obj Object) validateReflect(data any) Errors {
// 	val := reflect.ValueOf(data)
// 	kind := val.Kind()
// 	switch kind {
// 	case reflect.Struct:
// 		return obj.validateReflectStruct(val)
// 	case reflect.Map:
// 		panic("todo")
// 	case reflect.Pointer:
// 		panic("todo")
// 	case reflect.Slice:
// 		return newFieldError(obj.msg, "array")
// 	default:
// 		return newFieldError(obj.msg, kind.String())
// 	}
// }

func (obj Object) validateReflectStruct(data reflect.Value) Errors {
	panic("todo")
	// for i := range data.NumField() {
	// 	field := data.Field(i)
	// }
}

func (obj Object) Schema() jsony.Object {
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
		jsony.Field{K: "properties", V: properties},
	}
	if len(required) != 0 {
		res = append(res, jsony.Field{K: "required", V: required})
	}
	return res
}

type Property struct {
	name      string
	validator Validator
	optional  bool
}

func P(name string, v Validator) Property {
	return Property{name: name, validator: v}
}

func (p Property) Optional() Property {
	p.optional = true
	return p
}

func (p Property) Validate(data any) Error {
	err := p.validator.Validate(data)
	if err != nil {
		return ErrProperty{Name: p.name, Err: err}
	}
	return nil
}
