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
		return jsony.Float64(v)
	default:
		return jsony.Int(v)
	}
}

// The value must be a multiple of the given number.
//
// https://json-schema.org/understanding-json-schema/reference/numeric#multiples
func MultipleOf[T internal.Integer](v T) Constraint[T] {
	if v <= 0 {
		panic("the value must be positive")
	}
	c := func(f T) Error {
		// TODO: support float64 as well
		if f%v == 0 {
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

func Contains(v Validator) Constraint[[]any] {
	c := func(items []any) Error {
		var err Error
		for _, item := range items {
			err = v.Validate(item)
			if err == nil {
				return nil
			}
		}
		return ErrContains{Err: err}
	}
	return Constraint[[]any]{
		check: c,
		field: jsony.Field{K: "contains", V: v.Schema()},
	}
}

func MinItems(min uint) Constraint[[]any] {
	minInt := int(min)
	c := func(f []any) Error {
		if len(f) >= minInt {
			return nil
		}
		return ErrMinItems{Value: minInt}
	}
	return Constraint[[]any]{
		check: c,
		field: jsony.Field{K: "minItems", V: jsony.UInt(min)},
	}
}

func MaxItems(min uint) Constraint[[]any] {
	minInt := int(min)
	c := func(f []any) Error {
		if len(f) <= minInt {
			return nil
		}
		return ErrMaxItems{Value: minInt}
	}
	return Constraint[[]any]{
		check: c,
		field: jsony.Field{K: "maxItems", V: jsony.UInt(min)},
	}
}

func PropertyNames(cs ...Constraint[string]) Constraint[map[string]any] {
	c := func(items map[string]any) Error {
		res := Errors{}
		for _, c := range cs {
			for name := range items {
				err := c.check(name)
				if err != nil {
					res.Add(ErrPropertyNames{Name: name, Err: err})
				}
			}
		}
		return res.Flatten()
	}
	schema := make(jsony.Object, len(cs))
	for i, c := range cs {
		schema[i] = c.field
	}
	return Constraint[map[string]any]{
		check: c,
		field: jsony.Field{K: "propertyNames", V: schema},
	}
}

func MinProperties(min uint) Constraint[map[string]any] {
	minInt := int(min)
	c := func(f map[string]any) Error {
		if len(f) >= minInt {
			return nil
		}
		return ErrMinProperties{Value: minInt}
	}
	return Constraint[map[string]any]{
		check: c,
		field: jsony.Field{K: "minProperties", V: jsony.UInt(min)},
	}
}

func MaxProperties(min uint) Constraint[map[string]any] {
	minInt := int(min)
	c := func(f map[string]any) Error {
		if len(f) <= minInt {
			return nil
		}
		return ErrMaxProperties{Value: minInt}
	}
	return Constraint[map[string]any]{
		check: c,
		field: jsony.Field{K: "maxProperties", V: jsony.UInt(min)},
	}
}
