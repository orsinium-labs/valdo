package valdo

import "github.com/orsinium-labs/jsony"

type Meta struct {
	Validator   Validator
	Comment     string
	Title       string
	Description string
	Deprecated  bool
	// Default    any
	// Examples    []any
}

// Validate implements [Validator].
func (m Meta) Validate(data any) Error {
	return m.Validator.Validate(data)
}

// Schema implements [Validator].
func (m Meta) Schema() jsony.Object {
	s := m.Validator.Schema()
	if m.Comment != "" {
		s = append(s, jsony.Field{K: "$comment", V: jsony.String(m.Comment)})
	}
	if m.Title != "" {
		s = append(s, jsony.Field{K: "title", V: jsony.String(m.Title)})
	}
	if m.Description != "" {
		s = append(s, jsony.Field{K: "description", V: jsony.String(m.Description)})
	}
	if m.Deprecated {
		s = append(s, jsony.Field{K: "deprecated", V: jsony.True})
	}
	return s
}
