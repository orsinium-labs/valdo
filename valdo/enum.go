package valdo

import "github.com/orsinium-labs/jsony"

type enum struct {
	values []string
}

// Enum requires the input value to be one of the given constants.
//
// https://json-schema.org/understanding-json-schema/reference/enum
func Enum(vs ...string) Validator {
	return enum{vs}
}

// Validate implements [Validator].
func (v enum) Validate(data any) Error {
	got, err := stringValidator(data)
	if err != nil {
		return err
	}
	for _, exp := range v.values {
		if got == exp {
			return nil
		}
	}
	return ErrEnum{Got: got, Expected: v.values}
}

// Schema implements [Validator].
func (v enum) Schema() jsony.Object {
	values := make([]jsony.String, len(v.values))
	for i, val := range v.values {
		values[i] = jsony.String(val)
	}
	return jsony.Object{
		jsony.Field{K: "enum", V: jsony.Array[jsony.String](values)},
	}
}
