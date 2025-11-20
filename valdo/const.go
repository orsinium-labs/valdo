package valdo

import "github.com/orsinium-labs/jsony"

type strConst[T ~string] struct {
	value T
}

// Const restricts a value to a single value.
//
// https://json-schema.org/understanding-json-schema/reference/const
func StringConst[T ~string](value T) Validator {
	return strConst[T]{value}
}

// Validate implements [Validator].
func (v strConst[T]) Validate(data any) Error {
	got, ok := data.(T)
	if !ok {
		return ErrType{
			Got:      getTypeName(data),
			Expected: getTypeName(v.value),
		}
	}
	if got != v.value {
		return ErrConst{Got: got, Expected: v.value}
	}
	return nil
}

// Schema implements [Validator].
func (v strConst[T]) Schema() jsony.Object {
	return jsony.Object{
		jsony.Field{K: "const", V: jsony.String(v.value)},
	}
}

type boolConst[T ~bool] struct {
	value T
}

// Const restricts a value to a single value.
//
// https://json-schema.org/understanding-json-schema/reference/const
func BoolConst[T ~bool](value T) Validator {
	return boolConst[T]{value}
}

// Validate implements [Validator].
func (v boolConst[T]) Validate(data any) Error {
	got, ok := data.(T)
	if !ok {
		return ErrType{
			Got:      getTypeName(data),
			Expected: getTypeName(v.value),
		}
	}
	if got != v.value {
		return ErrConst{Got: got, Expected: v.value}
	}
	return nil
}

// Schema implements [Validator].
func (v boolConst[T]) Schema() jsony.Object {
	return jsony.Object{
		jsony.Field{K: "const", V: jsony.Bool(v.value)},
	}
}

type intConst[T ~int] struct {
	value T
}

// Const restricts a value to a single value.
//
// https://json-schema.org/understanding-json-schema/reference/const
func IntConst[T ~int](value T) Validator {
	return intConst[T]{value}
}

// Validate implements [Validator].
func (v intConst[T]) Validate(data any) Error {
	got, ok := data.(T)
	if !ok {
		return ErrType{
			Got:      getTypeName(data),
			Expected: getTypeName(v.value),
		}
	}
	if got != v.value {
		return ErrConst{Got: got, Expected: v.value}
	}
	return nil
}

// Schema implements [Validator].
func (v intConst[T]) Schema() jsony.Object {
	return jsony.Object{
		jsony.Field{K: "const", V: jsony.Int(v.value)},
	}
}
