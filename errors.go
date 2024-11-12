package valdo

import "fmt"

type Errors interface {
	Error() string
	First() *FieldError
}

type FieldError struct {
	Format string
	Args   []any
}

func newFieldError(f string, args ...any) *FieldError {
	return &FieldError{f, args}
}

func (e *FieldError) Error() string {
	return fmt.Sprintf(e.Format, e.Args...)
}
