package valdo

import (
	"fmt"
	"strings"
)

type Error interface {
	error
	GetDefault() Error
	SetFormat(f string) Error
}

type ErrorWrapper interface {
	Unwrap() error
	Map(func(Error) Error) Error
}

var (
	_ ErrorWrapper = ErrProperty{}
	_ ErrorWrapper = ErrIndex{}
	_ ErrorWrapper = ErrContains{}
	_ ErrorWrapper = ErrPropertyNames{}
)

type pair struct {
	name  string
	value any
}

func format(f string, pairs ...pair) string {
	args := make([]string, 0, len(pairs)*2)
	for _, p := range pairs {
		args = append(args, "{"+p.name+"}")
		args = append(args, fmt.Sprintf("%v", p.value))
	}
	return strings.NewReplacer(args...).Replace(f)
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

func (es Errors) GetDefault() Error {
	return Errors{}
}

func (es Errors) SetFormat(f string) Error {
	es.Sep = f
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
type ErrNoInput struct {
	Format string
}

func (e ErrNoInput) GetDefault() Error {
	return ErrNoInput{}
}

func (e ErrNoInput) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrNoInput) Error() string {
	f := e.Format
	if f == "" {
		f = "the input is empty"
	}
	return f
}

// An error in a field of an object.
type ErrProperty struct {
	Format string
	Name   string
	Err    Error
}

func (e ErrProperty) GetDefault() Error {
	return ErrProperty{}
}

func (e ErrProperty) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrProperty) Error() string {
	f := e.Format
	if f == "" {
		f = "{name}: {error}"
	}
	return format(f, pair{"name", e.Name}, pair{"error", e.Err})
}

func (e ErrProperty) Unwrap() error {
	return e.Err
}

func (e ErrProperty) Map(f func(Error) Error) Error {
	e.Err = f(e.Err)
	return e
}

// An error in an element of an array.
type ErrIndex struct {
	Format string
	Index  int
	Err    Error
}

func (e ErrIndex) GetDefault() Error {
	return ErrIndex{}
}

func (e ErrIndex) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrIndex) Error() string {
	f := e.Format
	if f == "" {
		f = "at {index}: {error}"
	}
	return format(f, pair{"index", e.Index}, pair{"error", e.Err})
}

func (e ErrIndex) Unwrap() error {
	return e.Err
}

func (e ErrIndex) Map(f func(Error) Error) Error {
	e.Err = f(e.Err)
	return e
}

// An error indicating a value of an unexpected type.
type ErrType struct {
	Format   string
	Got      string
	Expected string
}

func (e ErrType) GetDefault() Error {
	return ErrType{}
}

func (e ErrType) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrType) Error() string {
	f := e.Format
	if f == "" {
		f = "invalid type: got {got}, expected {expected}"
	}
	got := e.Got
	if got == "" {
		got = "unknown type"
	}
	return format(f, pair{"got", got}, pair{"expected", e.Expected})
}

// An error indicating that a value is required but not found.
type ErrRequired struct {
	Format string
	Name   string
}

func (e ErrRequired) GetDefault() Error {
	return ErrRequired{}
}

func (e ErrRequired) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrRequired) Error() string {
	f := e.Format
	if f == "" {
		f = "{name} is required but not found"
	}
	return format(f, pair{"name", e.Name})
}

// An error indicating that the property is not allowed.
type ErrUnexpected struct {
	Format string
	Name   string
}

func (e ErrUnexpected) GetDefault() Error {
	return ErrUnexpected{}
}

func (e ErrUnexpected) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrUnexpected) Error() string {
	f := e.Format
	if f == "" {
		f = "unexpected property: {name}"
	}
	return format(f, pair{"name", e.Name})
}

type ErrMultipleOf struct {
	Format string
	Value  any
}

func (e ErrMultipleOf) GetDefault() Error {
	return ErrMultipleOf{}
}

func (e ErrMultipleOf) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMultipleOf) Error() string {
	f := e.Format
	if f == "" {
		f = "must be a multiple of {value}"
	}
	return format(f, pair{"value", e.Value})
}

type ErrNot struct {
	Format string
}

func (e ErrNot) GetDefault() Error {
	return ErrNot{}
}

func (e ErrNot) SetFormat(f string) Error {
	e.Format = f
	return e
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

func (e ErrMin) GetDefault() Error {
	return ErrMin{}
}

func (e ErrMin) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMin) Error() string {
	f := e.Format
	if f == "" {
		f = "must be greater than or equal to {value}"
	}
	return format(f, pair{"value", e.Value})
}

type ErrExclMin struct {
	Format string
	Value  any
}

func (e ErrExclMin) GetDefault() Error {
	return ErrExclMin{}
}

func (e ErrExclMin) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrExclMin) Error() string {
	f := e.Format
	if f == "" {
		f = "must be greater than {value}"
	}
	return format(f, pair{"value", e.Value})
}

type ErrMax struct {
	Format string
	Value  any
}

func (e ErrMax) GetDefault() Error {
	return ErrMax{}
}

func (e ErrMax) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMax) Error() string {
	f := e.Format
	if f == "" {
		f = "must be less than or equal to {value}"
	}
	return format(f, pair{"value", e.Value})
}

type ErrExclMax struct {
	Format string
	Value  any
}

func (e ErrExclMax) GetDefault() Error {
	return ErrExclMax{}
}

func (e ErrExclMax) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrExclMax) Error() string {
	f := e.Format
	if f == "" {
		f = "must be less than {value}"
	}
	return format(f, pair{"value", e.Value})
}

type ErrMinLen struct {
	Format string
	Value  int
}

func (e ErrMinLen) GetDefault() Error {
	return ErrMinLen{}
}

func (e ErrMinLen) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMinLen) Error() string {
	f := e.Format
	if f == "" {
		f = "must be at least {value} characters long"
	}
	return format(f, pair{"value", e.Value})
}

type ErrMaxLen struct {
	Format string
	Value  int
}

func (e ErrMaxLen) GetDefault() Error {
	return ErrMaxLen{}
}

func (e ErrMaxLen) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMaxLen) Error() string {
	f := e.Format
	if f == "" {
		f = "must be at most {value} characters long"
	}
	return format(f, pair{"value", e.Value})
}

type ErrPattern struct {
	Format string
}

func (e ErrPattern) GetDefault() Error {
	return ErrPattern{}
}

func (e ErrPattern) SetFormat(f string) Error {
	e.Format = f
	return e
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
	Err    Error
}

func (e ErrContains) GetDefault() Error {
	return ErrContains{}
}

func (e ErrContains) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrContains) Error() string {
	f := e.Format
	if f == "" {
		f = "at least one item {error}"
	}
	return format(f, pair{"error", e.Err})
}

func (e ErrContains) Unwrap() error {
	return e.Err
}

func (e ErrContains) Map(f func(Error) Error) Error {
	e.Err = f(e.Err)
	return e
}

type ErrMinItems struct {
	Format string
	Value  int
}

func (e ErrMinItems) GetDefault() Error {
	return ErrMinItems{}
}

func (e ErrMinItems) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMinItems) Error() string {
	f := e.Format
	if f == "" {
		f = "must contain at least {value} items"
	}
	return format(f, pair{"value", e.Value})
}

type ErrMaxItems struct {
	Format string
	Value  int
}

func (e ErrMaxItems) GetDefault() Error {
	return ErrMaxItems{}
}

func (e ErrMaxItems) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMaxItems) Error() string {
	f := e.Format
	if f == "" {
		f = "must contain at most {value} items"
	}
	return format(f, pair{"value", e.Value})
}

type ErrPropertyNames struct {
	Format string
	Name   string
	Err    Error
}

func (e ErrPropertyNames) GetDefault() Error {
	return ErrPropertyNames{}
}

func (e ErrPropertyNames) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrPropertyNames) Error() string {
	f := e.Format
	if f == "" {
		f = "property name {name} {error}"
	}
	return format(f, pair{"name", e.Name}, pair{"error", e.Err})
}

func (e ErrPropertyNames) Unwrap() error {
	return e.Err
}

func (e ErrPropertyNames) Map(f func(Error) Error) Error {
	e.Err = f(e.Err)
	return e
}

type ErrMinProperties struct {
	Format string
	Value  int
}

func (e ErrMinProperties) GetDefault() Error {
	return ErrMinProperties{}
}

func (e ErrMinProperties) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMinProperties) Error() string {
	f := e.Format
	if f == "" {
		f = "must contain at least {value} properties"
	}
	return format(f, pair{"value", e.Value})
}

type ErrMaxProperties struct {
	Format string
	Value  int
}

func (e ErrMaxProperties) GetDefault() Error {
	return ErrMaxProperties{}
}

func (e ErrMaxProperties) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrMaxProperties) Error() string {
	f := e.Format
	if f == "" {
		f = "must contain at most {value} properties"
	}
	return format(f, pair{"value", e.Value})
}
