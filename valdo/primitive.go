package valdo

import (
	"math"

	"github.com/orsinium-labs/jsony"
	"github.com/orsinium-labs/valdo/internal"
)

// PrimitiveType is constructed by [Bool], [String], [Int], or [Float64].
type PrimitiveType[T internal.Primitive] struct {
	val  func(any) (T, Error)
	cs   []Constraint[T]
	name string
}

// Constrain adds constraints to the primitive, like [Min].
func (p PrimitiveType[T]) Constrain(cs ...Constraint[T]) PrimitiveType[T] {
	p.cs = append(p.cs, cs...)
	return p
}

// Validate implements [Validator].
func (p PrimitiveType[T]) Validate(raw any) Error {
	val, fErr := p.val(raw)
	if fErr != nil {
		return fErr
	}
	res := Errors{}
	for _, c := range p.cs {
		res.Add(c.check(val))
	}
	return res.Flatten()
}

// Schema implements [Validator].
func (p PrimitiveType[T]) Schema() jsony.Object {
	res := jsony.Object{
		jsony.Field{K: "type", V: jsony.String(p.name)},
	}
	for _, c := range p.cs {
		res = append(res, c.field)
	}
	return res
}

// Bool maps to "boolean" in JSON and "bool" in Go.
func Bool(cs ...Constraint[bool]) PrimitiveType[bool] {
	return PrimitiveType[bool]{
		val:  boolValidator,
		name: "boolean",
	}.Constrain(cs...)
}

func boolValidator(raw any) (bool, Error) {
	switch val := raw.(type) {
	case bool:
		return val, nil
	case jsony.Bool:
		return bool(val), nil
	case *bool:
		return *val, nil
	case *jsony.Bool:
		return bool(*val), nil
	default:
		return false, ErrType{Got: getTypeName(raw), Expected: "boolean"}
	}
}

// String maps to "string" in JSON and "string" in Go.
func String(cs ...Constraint[string]) PrimitiveType[string] {
	return PrimitiveType[string]{
		val:  stringValidator,
		name: "string",
	}.Constrain(cs...)
}

func stringValidator(raw any) (string, Error) {
	switch val := raw.(type) {
	case string:
		return val, nil
	case *string:
		return *val, nil
	case jsony.String:
		return string(val), nil
	case *jsony.String:
		return string(*val), nil
	default:
		return "", ErrType{Got: getTypeName(raw), Expected: "string"}
	}
}

// Int maps to "number" in JSON ("integer" in JSON Schema) and "int" in Go.
func Int(cs ...Constraint[int]) PrimitiveType[int] {
	return PrimitiveType[int]{
		val:  intValidator,
		name: "integer",
	}.Constrain(cs...)
}

func intValidator(raw any) (int, Error) {
	switch val := raw.(type) {
	case int:
		return val, nil
	case float64:
		if math.Floor(val) == val {
			return int(val), nil
		}
		return 0, ErrType{Got: "number", Expected: "integer"}
	case jsony.Int:
		return int(val), nil
	case *int:
		return *val, nil
	case *jsony.Int:
		return int(*val), nil
	default:
		return 0, ErrType{Got: getTypeName(raw), Expected: "integer"}
	}
}

// Float64 maps to "number" in JSON and "float64" in Go.
func Float64(cs ...Constraint[float64]) PrimitiveType[float64] {
	return PrimitiveType[float64]{
		val:  float64Validator,
		name: "number",
	}.Constrain(cs...)
}

func float64Validator(raw any) (float64, Error) {
	switch val := raw.(type) {
	case float64:
		return val, nil
	case jsony.Float64:
		return float64(val), nil
	case *float64:
		return *val, nil
	case *jsony.Float64:
		return float64(*val), nil
	case int:
		return float64(val), nil
	case *int:
		return float64(*val), nil
	default:
		return 0, ErrType{Got: getTypeName(raw), Expected: "number"}
	}
}

// nullType is constructed by [Null].
type nullType struct{}

// Null maps to "null" in JSON and "nil" in Go.
func Null() Validator {
	return nullType{}
}

// Validate implements [Validator].
func (n nullType) Validate(data any) Error {
	if data == nil {
		return nil
	}
	switch val := data.(type) {
	case []any:
		if val == nil {
			return nil
		}
	case map[string]any:
		if val == nil {
			return nil
		}
	}
	return ErrType{Got: getTypeName(data)}
}

// Schema implements [Validator].
func (n nullType) Schema() jsony.Object {
	return jsony.Object{
		jsony.Field{K: "type", V: jsony.SafeString("null")},
	}
}

type anyType struct{}

// Any is a value of any type.
func Any() Validator {
	return anyType{}
}

// Validate implements [Validator].
func (n anyType) Validate(data any) Error {
	return nil
}

// Schema implements [Validator].
func (n anyType) Schema() jsony.Object {
	return jsony.Object{}
}
