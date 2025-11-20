package valdo

import "github.com/orsinium-labs/jsony"

type TupleType struct {
	cs       []Constraint[[]any]
	vals     []Validator
	extra    bool
	extraVal Validator
}

var _ Validator = TupleType{}

// Validate array with every element having a different schema.
//
// https://json-schema.org/understanding-json-schema/reference/array#tupleValidation
func Tuple(vals ...Validator) TupleType {
	return TupleType{vals: vals}
}

func (t TupleType) Constrain(cs ...Constraint[[]any]) TupleType {
	t.cs = append(t.cs, cs...)
	return t
}

func (t TupleType) AllowExtra(v Validator) TupleType {
	t.extra = true
	t.extraVal = v
	return t
}

// Validate implements [Validator].
func (t TupleType) Validate(data any) Error {
	switch d := data.(type) {
	case []any:
		return t.validateArray(d)
	default:
		return ErrType{Got: getTypeName(data), Expected: "array"}
	}
}

func (t TupleType) validateArray(data []any) Error {
	if data == nil {
		return ErrType{Got: "null", Expected: "array"}
	}
	if len(data) < len(t.vals) {
		return ErrMinItems{Value: len(t.vals)}
	}
	if !t.extra && len(data) > len(t.vals) {
		return ErrMaxItems{Value: len(t.vals)}
	}
	res := Errors{}
	for i, validator := range t.vals {
		value := data[i]
		err := validator.Validate(value)
		if err != nil {
			res.Add(ErrIndex{Index: i, Err: err})
			break
		}
	}
	if t.extraVal != nil {
		for i := len(t.vals); i < len(data); i++ {
			value := data[i]
			err := t.extraVal.Validate(value)
			if err != nil {
				res.Add(ErrIndex{Index: i, Err: err})
				break
			}
		}
	}
	for _, c := range t.cs {
		res.Add(c.check(data))
	}
	return res.Flatten()
}

// Schema implements [Validator].
func (t TupleType) Schema() jsony.Object {
	var items jsony.Encoder
	if t.extraVal != nil {
		items = t.extraVal.Schema()
	} else {
		items = jsony.Bool(false)
	}
	res := jsony.Object{
		jsony.Field{K: "type", V: jsony.SafeString("array")},
		jsony.Field{K: "items", V: items},
	}
	if len(t.vals) > 0 {
		items := make([]jsony.Object, len(t.vals))
		for i, validator := range t.vals {
			items[i] = validator.Schema()
		}
		res = append(res, jsony.Field{K: "prefixItems", V: jsony.Array[jsony.Object](items)})
	}
	for _, c := range t.cs {
		res = append(res, c.field)
	}
	return res
}
