package valdo

import "github.com/orsinium-labs/jsony"

type allOf struct {
	vs []Validator
}

// AllOf requires all of the given validators to pass.
func AllOf(vs ...Validator) Validator {
	return allOf{vs: vs}
}

// Validate implements [Validator].
func (n allOf) Validate(data any) Error {
	for _, v := range n.vs {
		err := v.Validate(data)
		if err != nil {
			return err
		}
	}
	return nil
}

// Schema implements [Validator].
func (n allOf) Schema() jsony.Object {
	ss := make(jsony.Array[jsony.Object], len(n.vs))
	for i, v := range n.vs {
		ss[i] = v.Schema()
	}
	return jsony.Object{
		jsony.Field{K: "allOf", V: ss},
	}
}

type anyOf struct {
	vs []Validator
}

// Nullable allows null (nil) value for the given validator.
func Nullable(v Validator) Validator {
	return AnyOf(v, Null())
}

// AllOf requires at least one of the given validators to pass.
func AnyOf(vs ...Validator) Validator {
	return anyOf{vs: vs}
}

// Validate implements [Validator].
func (n anyOf) Validate(data any) Error {
	errors := Errors{}
	for _, v := range n.vs {
		err := v.Validate(data)
		if err == nil {
			return nil
		}
		errors.Add(err)
	}
	return ErrAnyOf{Errors: errors}
}

// Schema implements [Validator].
func (n anyOf) Schema() jsony.Object {
	ss := make(jsony.Array[jsony.Object], len(n.vs))
	for i, v := range n.vs {
		ss[i] = v.Schema()
	}
	return jsony.Object{
		jsony.Field{K: "anyOf", V: ss},
	}
}

type notType struct {
	v Validator
}

func Not(v Validator) Validator {
	return notType{v: v}
}

// Validate implements [Validator].
func (n notType) Validate(data any) Error {
	err := n.v.Validate(data)
	if err == nil {
		return ErrNot{}
	}
	return nil
}

// Schema implements [Validator].
func (n notType) Schema() jsony.Object {
	return jsony.Object{
		jsony.Field{K: "not", V: n.v.Schema()},
	}
}
