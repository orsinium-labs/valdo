package valdo

import (
	"fmt"
	"strings"
)

type Error = error

var (
	_ Error = Errors{}
	_ Error = ErrProperty{}
	_ Error = ErrIndex{}
	_ Error = ErrType{}
	_ Error = ErrRequired{}
	_ Error = ErrUnexpected{}
)

// A collection of multiple errors.
type Errors struct {
	Sep  string
	Errs []Error
}

func (es Errors) Flatten() Error {
	if len(es.Errs) == 0 {
		return nil
	}
	if len(es.Errs) == 1 {
		return es.Errs[0]
	}
	return es
}

func (es Errors) Error() string {
	res := make([]string, len(es.Errs))
	for i, e := range es.Errs {
		res[i] = e.Error()
	}
	return strings.Join(res, es.Sep)
}

// An error in a field of an object.
type ErrProperty struct {
	Format string
	Name   string
	Err    Error
}

func (e ErrProperty) Error() string {
	f := e.Format
	if f == "" {
		f = "%s: %v"
	}
	return fmt.Sprintf(f, e.Name, e.Err)
}

// An error in an element of an array.
type ErrIndex struct {
	Format string
	Index  int
	Err    Error
}

func (e ErrIndex) Error() string {
	f := e.Format
	if f == "" {
		f = "at %d: %v"
	}
	return fmt.Sprintf(f, e.Index, e.Err)
}

// An error indicating a value of an unexpected type.
type ErrType struct {
	Format   string
	Got      string
	Expected string
}

func (e ErrType) Error() string {
	f := e.Format
	if f == "" {
		f = "invalid type: got %s, expected %s"
	}
	return fmt.Sprintf(f, e.Got, e.Expected)
}

// An error indicating that a value is required but not found.
type ErrRequired struct {
	Format string
	Name   string
}

func (e ErrRequired) Error() string {
	f := e.Format
	if f == "" {
		f = "%s is required but not found"
	}
	return fmt.Sprintf(f, e.Name)
}

// An error indicating that the property is not allowed.
type ErrUnexpected struct {
	Format string
	Name   string
}

func (e ErrUnexpected) Error() string {
	f := e.Format
	if f == "" {
		f = "unexpected property: %s"
	}
	return fmt.Sprintf(f, e.Name)
}

type ErrMultipleOf struct {
	Format string
	Value  any
}

func (e ErrMultipleOf) Error() string {
	f := e.Format
	if f == "" {
		f = "must be a multiple of %v"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrMinimum struct {
	Format string
	Value  any
}

func (e ErrMinimum) Error() string {
	f := e.Format
	if f == "" {
		f = "must be greater than or equal to %v"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrMinLength struct {
	Format string
	Value  int
}

func (e ErrMinLength) Error() string {
	f := e.Format
	if f == "" {
		f = "must be at least %d characters long"
	}
	return fmt.Sprintf(f, e.Value)
}
