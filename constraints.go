package valdo

import (
	"github.com/orsinium-labs/jsony"
	"github.com/orsinium-labs/valdo/internal"
)

type Constraint[T any] struct {
	check func(T) Error
	field jsony.Field
}

func jsonyNumber[T internal.Number](v T) jsony.Encoder {
	switch any(v).(type) {
	case float32:
		return jsony.Float32(v)
	case float64:
		return jsony.Float32(v)
	default:
		return jsony.I(v)
	}
}

// A numeric instance is valid only if division by this keyword's value results in an integer.
//
// https://json-schema.org/draft/2020-12/json-schema-validation#name-multipleof
func MultipleOf[T internal.Number](v T) Constraint[T] {
	if v <= 0 {
		panic("the value must be positive")
	}
	c := func(f T) Error {
		if f/v == 0 {
			return nil
		}
		return ErrMultipleOf{Value: v}
	}
	return Constraint[T]{
		check: c,
		field: jsony.Field{K: "multipleOf", V: jsonyNumber(v)},
	}
}

func Minimum[T internal.Number](v T) Constraint[T] {
	c := func(f T) Error {
		if f >= v {
			return nil
		}
		return ErrMinimum{Value: v}
	}
	return Constraint[T]{
		check: c,
		field: jsony.Field{K: "minimum", V: jsonyNumber(v)},
	}
}

func MinLength(min uint) Constraint[string] {
	minInt := int(min)
	c := func(f string) Error {
		if len(f) >= minInt {
			return nil
		}
		return ErrMinLength{Value: minInt}
	}
	return Constraint[string]{
		check: c,
		field: jsony.Field{K: "minLength", V: jsony.UInt(min)},
	}
}
