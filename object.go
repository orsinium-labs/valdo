package valdo

import "github.com/orsinium-labs/jsony"

type Object struct {
	ps []Property
}

var _ Validator = O()

func O(ps ...Property) Object {
	return Object{
		ps: ps,
	}
}

func (obj Object) Validate(data any) Errors {
	panic("todo")
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
