package valdo

import "github.com/orsinium-labs/jsony"

type constVal[T string | bool | int] struct {
	validator func(any) (T, Error)
	value     T
}

func (p constVal[T]) Validate(raw any) Error {
	got, fErr := p.validator(raw)
	if fErr != nil {
		return fErr
	}
	if got != p.value {
		return ErrConst{Got: got, Expected: p.value}
	}
	return nil
}

// Schema implements [Validator].
func (p constVal[T]) Schema() jsony.Object {
	field := jsony.Field{K: "const", V: jsony.Detect(p.value)}
	return jsony.Object{field}
}

// DEPRECATED: Use [StringConst] instead.
func Const[T ~string](value T) Validator {
	return StringConst(string(value))
}

// StringConst restricts a value to a single string value.
//
// https://json-schema.org/understanding-json-schema/reference/const
func StringConst(value string) Validator {
	return constVal[string]{
		validator: stringValidator,
		value:     value,
	}
}

// BoolConst restricts a value to a single boolean value.
//
// https://json-schema.org/understanding-json-schema/reference/const
func BoolConst(value bool) Validator {
	return constVal[bool]{
		validator: boolValidator,
		value:     value,
	}

}

// IntConst restricts a value to a single integer value.
//
// https://json-schema.org/understanding-json-schema/reference/const
func IntConst(value int) Validator {
	return constVal[int]{
		validator: intValidator,
		value:     value,
	}
}
