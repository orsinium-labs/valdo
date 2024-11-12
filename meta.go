package valdo

import "github.com/orsinium-labs/jsony"

type Meta struct {
	field jsony.Field
}

func Title(t string) Meta {
	return Meta{jsony.Field{K: "title", V: jsony.String(t)}}
}

func Description(t string) Meta {
	return Meta{jsony.Field{K: "description", V: jsony.String(t)}}
}

func Deprecated() Meta {
	return Meta{jsony.Field{K: "deprecated", V: jsony.True}}
}
