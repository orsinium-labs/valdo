package valdo

import (
	"github.com/orsinium-labs/jsony"
)

type Validator interface {
	Validate(data any) Errors
	Schema() jsony.Object
}

type Constraint[T any] struct {
	check   func(string, T) *FieldError
	field   jsony.Field
	message string
}

// Make a copy of the check with the given error message.
func (f Constraint[T]) M(message string) Constraint[T] {
	f.message = message
	return f
}
