package valdo

import (
	"reflect"

	"github.com/orsinium-labs/jsony"
)

type Object struct {
	ps    []Property
	msg   string
	extra bool
}

var _ Validator = O()

func O(ps ...Property) Object {
	return Object{
		ps:  ps,
		msg: "invalid type: expected object, got %s",
	}
}

func (obj Object) Validate(data any) Errors {
	switch d := data.(type) {
	case map[string]any:
		return obj.validateMap(d)
	case int, int8, int16, int32, int64:
		return newFieldError(obj.msg, "integer")
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return newFieldError(obj.msg, "unsigned integer")
	case bool:
		return newFieldError(obj.msg, "boolean")
	case string:
		return newFieldError(obj.msg, "string")
	case float32, float64:
		return newFieldError(obj.msg, "number")
	default:
		return obj.validateReflect(data)
	}
}

func (obj Object) validateMap(data map[string]any) Errors {
	var resErr Errors
	for _, p := range obj.ps {
		val, found := data[p.name]
		if !found {
			if p.optional {
				continue
			}
			resErr = newFieldError("%s required but not found", p.name)
		}
		pErr := p.Validate(val)
		if pErr != nil {
			// TODO: concatenate errors
			resErr = pErr
		}
	}
	if !obj.extra {
		for name := range data {
			if !obj.hasProperty(name) {
				return newFieldError("unexpected property found: %s", name)
			}
		}
	}
	return resErr
}

func (obj Object) hasProperty(name string) bool {
	for _, p := range obj.ps {
		if p.name == name {
			return true
		}
	}
	return false
}

func (obj Object) validateReflect(data any) Errors {
	val := reflect.ValueOf(data)
	kind := val.Kind()
	switch kind {
	case reflect.Struct:
		return obj.validateReflectStruct(val)
	case reflect.Map:
		panic("todo")
	case reflect.Pointer:
		panic("todo")
	case reflect.Slice:
		return newFieldError(obj.msg, "array")
	default:
		return newFieldError(obj.msg, kind.String())
	}
}

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

func (p Property) Validate(data any) Errors {
	err := p.validator.Validate(data)
	if err != nil {
		return newFieldError(p.name+" %v", err.Error())
	}
	return nil
}
