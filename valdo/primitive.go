package valdo

import (
	"github.com/orsinium-labs/jsony"
	"github.com/orsinium-labs/valdo/internal"
)

type PrimitiveType[T internal.Primitive] struct {
	val  func(any) (T, Error)
	cs   []Constraint[T]
	name string
}

func (p PrimitiveType[T]) Constrain(cs ...Constraint[T]) PrimitiveType[T] {
	p.cs = append(p.cs, cs...)
	return p
}

func (p PrimitiveType[T]) Validate(raw any) Error {
	val, fErr := p.val(raw)
	if fErr != nil {
		return fErr
	}
	for _, c := range p.cs {
		fErr := c.check(val)
		if fErr != nil {
			return fErr
		}
	}
	return nil
}

func (p PrimitiveType[T]) Schema() jsony.Object {
	res := jsony.Object{
		jsony.Field{K: "type", V: jsony.String(p.name)},
	}
	for _, c := range p.cs {
		res = append(res, c.field)
	}
	return res
}

func Bool() PrimitiveType[bool] {
	return PrimitiveType[bool]{
		val:  boolValidator,
		name: "boolean",
	}
}

func boolValidator(raw any) (bool, Error) {
	switch val := raw.(type) {
	case bool:
		return val, nil
	case jsony.Bool:
		return bool(val), nil
	default:
		return false, ErrType{Got: getTypeName(raw), Expected: "boolean"}
	}
}

func String() PrimitiveType[string] {
	return PrimitiveType[string]{
		val:  stringValidator,
		name: "string",
	}
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

func Int() PrimitiveType[int] {
	return PrimitiveType[int]{
		val:  intValidator,
		name: "integer",
	}
}

func intValidator(raw any) (int, Error) {
	switch val := raw.(type) {
	case int:
		return val, nil
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

func Float64() PrimitiveType[float64] {
	return PrimitiveType[float64]{
		val:  float64Validator,
		name: "number",
	}
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
