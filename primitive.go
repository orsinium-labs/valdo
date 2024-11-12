package valdo

import (
	"github.com/orsinium-labs/jsony"
	"github.com/orsinium-labs/valdo/internal"
)

type Primitive[T internal.Primitive] struct {
	val  func(string, any) (T, *FieldError)
	cs   []Constraint[T]
	meta []Meta
	name string
	msg  string
}

var (
	_ Validator = Bool()
	_ Validator = String()
)

func (p Primitive[T]) Message(msg string) Primitive[T] {
	p.msg = msg
	return p
}

func (p Primitive[T]) Meta(meta ...Meta) Primitive[T] {
	p.meta = append(p.meta, meta...)
	return p
}

func (p Primitive[T]) Constrain(cs ...Constraint[T]) Primitive[T] {
	p.cs = append(p.cs, cs...)
	return p
}

func (p Primitive[T]) Validate(raw any) Errors {
	val, fErr := p.val(p.msg, raw)
	if fErr != nil {
		return fErr
	}
	for _, c := range p.cs {
		fErr := c.check(c.message, val)
		if fErr != nil {
			return fErr
		}
	}
	return nil
}

func (p Primitive[T]) Schema() jsony.Object {
	res := jsony.Object{
		jsony.Field{K: "type", V: jsony.String(p.name)},
	}
	for _, c := range p.cs {
		res = append(res, c.field)
	}
	for _, meta := range p.meta {
		res = append(res, meta.field)
	}
	return res
}

func Bool() Primitive[bool] {
	return Primitive[bool]{
		val:  boolValidator,
		msg:  "must be boolean",
		name: "boolean",
	}
}

func boolValidator(msg string, raw any) (bool, *FieldError) {
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
		return false, newFieldError(msg)
	}
	return false, newFieldError(msg)
}

func String() Primitive[string] {
	return Primitive[string]{
		val:  stringValidator,
		msg:  "must be string",
		name: "string",
	}
}

func stringValidator(msg string, raw any) (string, *FieldError) {
	val, isString := raw.(string)
	if !isString {
		return "", newFieldError(msg)
	}
	return val, nil
}
