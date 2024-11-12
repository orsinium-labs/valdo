package valdo

import (
	"github.com/orsinium-labs/jsony"
	"github.com/orsinium-labs/valdo/internal"
)

type Constraint[T any] struct {
	check func(string, T) *FieldError
	field jsony.Field
	msg   string
}

func (c Constraint[T]) Message(msg string) Constraint[T] {
	c.msg = msg
	return c
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
	c := func(m string, f T) *FieldError {
		if f/v == 0 {
			return nil
		}
		return newFieldError(m, v)
	}
	return Constraint[T]{
		check: c,
		field: jsony.Field{K: "multipleOf", V: jsonyNumber(v)},
		msg:   "must be a multiple of %v",
	}
}

func Minimum[T internal.Number](v T) Constraint[T] {
	c := func(m string, f T) *FieldError {
		if f >= v {
			return nil
		}
		return newFieldError(m, v)
	}
	return Constraint[T]{
		check: c,
		field: jsony.Field{K: "minimum", V: jsonyNumber(v)},
		msg:   "must be greater than or equal to %v",
	}
}
