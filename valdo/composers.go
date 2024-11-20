package valdo

import "github.com/orsinium-labs/jsony"

type allOf struct {
	vs []Validator
}

func AllOf(vs ...Validator) Validator {
	return allOf{vs: vs}
}

func (n allOf) Validate(data any) Error {
	for _, v := range n.vs {
		err := v.Validate(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n allOf) Schema() jsony.Object {
	ss := make(jsony.Array[jsony.Object], len(n.vs))
	for i, v := range n.vs {
		ss[i] = v.Schema()
	}
	return jsony.Object{
		jsony.Field{K: "allOf", V: ss},
	}
}

type notType struct {
	v Validator
}

func Not(v Validator) Validator {
	return notType{v: v}
}

func (n notType) Validate(data any) Error {
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
