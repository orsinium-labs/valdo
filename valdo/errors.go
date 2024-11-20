package valdo

import (
	"fmt"
	"strings"
)

type Error interface {
	error
}

// A collection of multiple errors.
type Errors struct {
	Sep  string
	Errs []Error
}

// Add the given error (if not nil) to the list of errors.
func (es *Errors) Add(err Error) {
	if err != nil {
		es.Errs = append(es.Errs, err)
	}
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

// Error implements [error] interface.
func (es Errors) Error() string {
	res := make([]string, len(es.Errs))
	for i, e := range es.Errs {
		res[i] = e.Error()
	}
	sep := es.Sep
	if sep == "" {
		sep = "; "
	}
	return strings.Join(res, sep)
}

// Unwrap makes errors.Is work.
func (es Errors) Unwrap() []error {
	res := make([]error, len(es.Errs))
	for i, subErr := range es.Errs {
		res[i] = subErr
	}
	return res
}

// An error in a field of an object.
type ErrProperty struct {
	Format string
	Name   string
	Err    Error
}

// Error implements [error] interface.
func (e ErrProperty) Error() string {
	f := e.Format
	if f == "" {
		f = "%s: %v"
	}
	return fmt.Sprintf(f, e.Name, e.Err)
}

func (e ErrProperty) Unwrap() error {
	return e.Err
}

// An error in an element of an array.
type ErrIndex struct {
	Format string
	Index  int
	Err    Error
}

// Error implements [error] interface.
func (e ErrIndex) Error() string {
	f := e.Format
	if f == "" {
		f = "at %d: %v"
	}
	return fmt.Sprintf(f, e.Index, e.Err)
}

func (e ErrIndex) Unwrap() error {
	return e.Err
}

// An error indicating a value of an unexpected type.
type ErrType struct {
	Format   string
	Got      string
	Expected string
}

// Error implements [error] interface.
func (e ErrType) Error() string {
	f := e.Format
	if f == "" {
		f = "invalid type: got %s, expected %s"
	}
	got := e.Got
	if got == "" {
		got = "unknown type"
	}
	return fmt.Sprintf(f, got, e.Expected)
}

// An error indicating that a value is required but not found.
type ErrRequired struct {
	Format string
	Name   string
}

// Error implements [error] interface.
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

// Error implements [error] interface.
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

// Error implements [error] interface.
func (e ErrMultipleOf) Error() string {
	f := e.Format
	if f == "" {
		f = "must be a multiple of %v"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrNot struct {
	Format string
}

// Error implements [error] interface.
func (e ErrNot) Error() string {
	f := e.Format
	if f == "" {
		f = "must not match the schema"
	}
	return f
}

type ErrMin struct {
	Format string
	Value  any
}

// Error implements [error] interface.
func (e ErrMin) Error() string {
	f := e.Format
	if f == "" {
		f = "must be greater than or equal to %v"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrExclMin struct {
	Format string
	Value  any
}

// Error implements [error] interface.
func (e ErrExclMin) Error() string {
	f := e.Format
	if f == "" {
		f = "must be greater than %v"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrMax struct {
	Format string
	Value  any
}

// Error implements [error] interface.
func (e ErrMax) Error() string {
	f := e.Format
	if f == "" {
		f = "must be less than or equal to %v"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrExclMax struct {
	Format string
	Value  any
}

// Error implements [error] interface.
func (e ErrExclMax) Error() string {
	f := e.Format
	if f == "" {
		f = "must be less than %v"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrMinLen struct {
	Format string
	Value  int
}

// Error implements [error] interface.
func (e ErrMinLen) Error() string {
	f := e.Format
	if f == "" {
		f = "must be at least %d characters long"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrMaxLen struct {
	Format string
	Value  int
}

// Error implements [error] interface.
func (e ErrMaxLen) Error() string {
	f := e.Format
	if f == "" {
		f = "must be at most %d characters long"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrPattern struct {
	Format string
}

// Error implements [error] interface.
func (e ErrPattern) Error() string {
	f := e.Format
	if f == "" {
		f = "must match the pattern"
	}
	return f
}

type ErrContains struct {
	Format string
	Err    error
}

// Error implements [error] interface.
func (e ErrContains) Error() string {
	f := e.Format
	if f == "" {
		f = "at least one item %v"
	}
	return fmt.Sprintf(f, e.Err)
}

func (e ErrContains) Unwrap() error {
	return e.Err
}

type ErrMinItems struct {
	Format string
	Value  int
}

// Error implements [error] interface.
func (e ErrMinItems) Error() string {
	f := e.Format
	if f == "" {
		f = "must contain at least %d items"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrMaxItems struct {
	Format string
	Value  int
}

// Error implements [error] interface.
func (e ErrMaxItems) Error() string {
	f := e.Format
	if f == "" {
		f = "must contain at most %d items"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrPropertyNames struct {
	Format string
	Name   string
	Err    error
}

// Error implements [error] interface.
func (e ErrPropertyNames) Error() string {
	f := e.Format
	if f == "" {
		f = "property name %s %v"
	}
	return fmt.Sprintf(f, e.Name, e.Err)
}

func (e ErrPropertyNames) Unwrap() error {
	return e.Err
}

type ErrMinProperties struct {
	Format string
	Value  int
}

// Error implements [error] interface.
func (e ErrMinProperties) Error() string {
	f := e.Format
	if f == "" {
		f = "must contain at least %d properties"
	}
	return fmt.Sprintf(f, e.Value)
}

type ErrMaxProperties struct {
	Format string
	Value  int
}

// Error implements [error] interface.
func (e ErrMaxProperties) Error() string {
	f := e.Format
	if f == "" {
		f = "must contain at most %d properties"
	}
	return fmt.Sprintf(f, e.Value)
}
