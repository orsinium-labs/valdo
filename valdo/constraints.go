package valdo

import (
	"regexp"

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

func Min[T internal.Number](v T) Constraint[T] {
	c := func(f T) Error {
		if f >= v {
			return nil
		}
		return ErrMin{Value: v}
	}
	return Constraint[T]{
		check: c,
		field: jsony.Field{K: "minimum", V: jsonyNumber(v)},
	}
}

func ExclMin[T internal.Number](v T) Constraint[T] {
	c := func(f T) Error {
		if f > v {
			return nil
		}
		return ErrExclMin{Value: v}
	}
	return Constraint[T]{
		check: c,
		field: jsony.Field{K: "exclusiveMinimum", V: jsonyNumber(v)},
	}
}

func Max[T internal.Number](v T) Constraint[T] {
	c := func(f T) Error {
		if f <= v {
			return nil
		}
		return ErrMax{Value: v}
	}
	return Constraint[T]{
		check: c,
		field: jsony.Field{K: "maximum", V: jsonyNumber(v)},
	}
}

func ExclMax[T internal.Number](v T) Constraint[T] {
	c := func(f T) Error {
		if f < v {
			return nil
		}
		return ErrExclMax{Value: v}
	}
	return Constraint[T]{
		check: c,
		field: jsony.Field{K: "exclusiveMaximum", V: jsonyNumber(v)},
	}
}

func MinLen(min uint) Constraint[string] {
	minInt := int(min)
	c := func(f string) Error {
		if len(f) >= minInt {
			return nil
		}
		return ErrMinLen{Value: minInt}
	}
	return Constraint[string]{
		check: c,
		field: jsony.Field{K: "minLength", V: jsony.UInt(min)},
	}
}

func MaxLen(min uint) Constraint[string] {
	minInt := int(min)
	c := func(f string) Error {
		if len(f) <= minInt {
			return nil
		}
		return ErrMaxLen{Value: minInt}
	}
	return Constraint[string]{
		check: c,
		field: jsony.Field{K: "maxLength", V: jsony.UInt(min)},
	}
}

func Pattern(r string) Constraint[string] {
	rex := regexp.MustCompile(r)
	c := func(f string) Error {
		if rex.MatchString(f) {
			return nil
		}
		return ErrPattern{}
	}
	return Constraint[string]{
		check: c,
		field: jsony.Field{K: "pattern", V: jsony.String(r)},
	}
}
