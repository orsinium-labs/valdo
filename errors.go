package valdo

import "fmt"

type Errors interface {
	Error() string
	First() *FieldError
}

type FieldError struct {
	Template string
	Args     []any
}

func newFieldError(t string, args ...any) *FieldError {
	return &FieldError{t, args}
}

func (e *FieldError) Error() string {
	return fmt.Sprintf(e.Template, e.Args...)
}
