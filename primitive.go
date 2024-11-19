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

var (
	_ Validator = Bool()
	_ Validator = String()
)

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
	val, isBool := raw.(bool)
	if !isBool {
		return false, ErrType{Got: getTypeName(raw), Expected: "boolean"}
	}
	return val, nil
}

func String() PrimitiveType[string] {
	return PrimitiveType[string]{
		val:  stringValidator,
		name: "string",
	}
}

func stringValidator(raw any) (string, Error) {
	val, isString := raw.(string)
	if !isString {
		return "", ErrType{Got: getTypeName(raw), Expected: "string"}
	}
	return val, nil
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
		return int(val), nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case uintptr:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	default:
		return 0, ErrType{Got: getTypeName(raw), Expected: "integer"}
	}
}
