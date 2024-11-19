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
	switch val := raw.(type) {
	case bool:
		return val, nil
	case string:
		switch val {
		case "true":
			return true, nil
		case "false":
			return false, nil
		}
		return false, ErrType{Got: "string", Expected: "boolean"}
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
	val, isString := raw.(string)
	if !isString {
		return "", ErrType{Got: getTypeName(raw), Expected: "string"}
	}
	return val, nil
}
