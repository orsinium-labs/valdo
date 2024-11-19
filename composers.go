package valdo

import "github.com/orsinium-labs/jsony"

type notType struct {
	v Validator
}

func Not(v Validator) Validator {
	return notType{v: v}
}

func (n notType) Validate(data any) error {
	err := n.v.Validate(data)
	if err == nil {
		return ErrNot{}
	}
	return nil
}

func (n notType) Schema() jsony.Object {
	return jsony.Object{
		jsony.Field{K: "not", V: n.v.Schema()},
	}
}
