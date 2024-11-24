package valdo

import (
	"fmt"
	"strings"
)

type Error interface {
	error
	// GetDefault implements [Error] interface.
	GetDefault() Error
	// SetFormat implements [Error] interface.
	SetFormat(f string) Error
}

// ErrorWrapper is an [Error] that wraps another error.
type ErrorWrapper interface {
	Error
	// Unwrap makes it possible for errors.Unwrap function to access the wrapped error.
	Unwrap() error

	// Map applies the given function to the inner function and returns it, wrapped.
	//
	// Yes, it's a monad! In Go! What a day.
	Map(func(Error) Error) Error
}

var (
	_ ErrorWrapper = ErrProperty{}
	_ ErrorWrapper = ErrIndex{}
	_ ErrorWrapper = ErrContains{}
	_ ErrorWrapper = ErrPropertyNames{}
	_ ErrorWrapper = ErrAnyOf{}
)

type pair struct {
	name  string
	value any
}

// Format substitutes values into a format string with python-style placeholders.
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

// GetDefault implements [Error] interface.
func (es Errors) GetDefault() Error {
	return Errors{}
}

// SetFormat implements [Error] interface.
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

// Map applies the given function to the inner function and returns it, wrapped.
//
// It partially implements [ErrorWrapper]: Errors doesn't satisfy the interface
// because [Errors.Unwrap] returns a list of errrors instead of a single error.
func (e Errors) Map(f func(Error) Error) Error {
	errors := make([]Error, len(e.Errs))
	for i, sub := range e.Errs {
		errors[i] = f(sub)
	}
	e.Errs = errors
	return e
}

// Unwrap makes errors.Is work.
//
// Note that since it returns a list of errors instead of a single error,
// errors.Unwrap will NOT work.
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

// GetDefault implements [Error] interface.
func (e ErrNoInput) GetDefault() Error {
	return ErrNoInput{}
}

// SetFormat implements [Error] interface.
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
//
// Returned by an [Object] validator.
type ErrProperty struct {
	Format string
	Name   string
	Err    Error
}

// GetDefault implements [Error] interface.
func (e ErrProperty) GetDefault() Error {
	return ErrProperty{}
}

// SetFormat implements [Error] interface.
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

// Unwrap implements [ErrorWrapper] interface.
func (e ErrProperty) Unwrap() error {
	return e.Err
}

// Map implements [ErrorWrapper] interface.
func (e ErrProperty) Map(f func(Error) Error) Error {
	e.Err = f(e.Err)
	return e
}

// An error in an element of an array.
//
// Returned by an [Array] validator.
type ErrIndex struct {
	Format string
	Index  int
	Err    Error
}

// GetDefault implements [Error] interface.
func (e ErrIndex) GetDefault() Error {
	return ErrIndex{}
}

// SetFormat implements [Error] interface.
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

// Unwrap implements [ErrorWrapper] interface.
func (e ErrIndex) Unwrap() error {
	return e.Err
}

// Map implements [ErrorWrapper] interface.
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

// GetDefault implements [Error] interface.
func (e ErrType) GetDefault() Error {
	return ErrType{}
}

// SetFormat implements [Error] interface.
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
//
// Returned by an [Object] validator.
type ErrRequired struct {
	Format string
	Name   string
}

// GetDefault implements [Error] interface.
func (e ErrRequired) GetDefault() Error {
	return ErrRequired{}
}

// SetFormat implements [Error] interface.
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
//
// Returned by an [Object] validator.
type ErrUnexpected struct {
	Format string
	Name   string
}

// GetDefault implements [Error] interface.
func (e ErrUnexpected) GetDefault() Error {
	return ErrUnexpected{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [MultipleOf].
type ErrMultipleOf struct {
	Format string
	Value  any
}

// GetDefault implements [Error] interface.
func (e ErrMultipleOf) GetDefault() Error {
	return ErrMultipleOf{}
}

// SetFormat implements [Error] interface.
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

// An error returned by [Not] validator.
type ErrNot struct {
	Format string
}

// GetDefault implements [Error] interface.
func (e ErrNot) GetDefault() Error {
	return ErrNot{}
}

// SetFormat implements [Error] interface.
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

// An error returned by [AnyOf] validator.
type ErrAnyOf struct {
	Format string
	Errors Errors
}

// Map implements [ErrorWrapper] interface.
func (e ErrAnyOf) Map(f func(Error) Error) Error {
	errors := make([]Error, len(e.Errors.Errs))
	for i, sub := range e.Errors.Errs {
		errors[i] = f(sub)
	}
	e.Errors = Errors{Errs: errors}
	return e
}

// Unwrap implements [ErrorWrapper] interface.
func (e ErrAnyOf) Unwrap() error {
	return e.Errors
}

// GetDefault implements [Error] interface.
func (e ErrAnyOf) GetDefault() Error {
	return ErrAnyOf{}
}

// SetFormat implements [Error] interface.
func (e ErrAnyOf) SetFormat(f string) Error {
	e.Format = f
	return e
}

// Error implements [error] interface.
func (e ErrAnyOf) Error() string {
	f := e.Format
	if f == "" {
		f = "must match any of the conditions: {errors}"
	}
	return format(f, pair{"errors", e.Errors})
}

// A constraint error returned by [Min].
type ErrMin struct {
	Format string
	Value  any
}

// GetDefault implements [Error] interface.
func (e ErrMin) GetDefault() Error {
	return ErrMin{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [ExclMin].
type ErrExclMin struct {
	Format string
	Value  any
}

// GetDefault implements [Error] interface.
func (e ErrExclMin) GetDefault() Error {
	return ErrExclMin{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [Max].
type ErrMax struct {
	Format string
	Value  any
}

// GetDefault implements [Error] interface.
func (e ErrMax) GetDefault() Error {
	return ErrMax{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [ExclMax].
type ErrExclMax struct {
	Format string
	Value  any
}

// GetDefault implements [Error] interface.
func (e ErrExclMax) GetDefault() Error {
	return ErrExclMax{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [MinLen].
type ErrMinLen struct {
	Format string
	Value  int
}

// GetDefault implements [Error] interface.
func (e ErrMinLen) GetDefault() Error {
	return ErrMinLen{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [MaxLen].
type ErrMaxLen struct {
	Format string
	Value  int
}

// GetDefault implements [Error] interface.
func (e ErrMaxLen) GetDefault() Error {
	return ErrMaxLen{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [Pattern].
type ErrPattern struct {
	Format string
}

// GetDefault implements [Error] interface.
func (e ErrPattern) GetDefault() Error {
	return ErrPattern{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [Contains].
type ErrContains struct {
	Format string
	Err    Error
}

// GetDefault implements [Error] interface.
func (e ErrContains) GetDefault() Error {
	return ErrContains{}
}

// SetFormat implements [Error] interface.
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

// Unwrap implements [ErrorWrapper] interface.
func (e ErrContains) Unwrap() error {
	return e.Err
}

// Map implements [ErrorWrapper] interface.
func (e ErrContains) Map(f func(Error) Error) Error {
	e.Err = f(e.Err)
	return e
}

// A constraint error returned by [MinItems].
type ErrMinItems struct {
	Format string
	Value  int
}

// GetDefault implements [Error] interface.
func (e ErrMinItems) GetDefault() Error {
	return ErrMinItems{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [MaxItems].
type ErrMaxItems struct {
	Format string
	Value  int
}

// GetDefault implements [Error] interface.
func (e ErrMaxItems) GetDefault() Error {
	return ErrMaxItems{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [PropertyNames].
type ErrPropertyNames struct {
	Format string
	Name   string
	Err    Error
}

// GetDefault implements [Error] interface.
func (e ErrPropertyNames) GetDefault() Error {
	return ErrPropertyNames{}
}

// SetFormat implements [Error] interface.
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

// Unwrap implements [ErrorWrapper] interface.
func (e ErrPropertyNames) Unwrap() error {
	return e.Err
}

// Map implements [ErrorWrapper] interface.
func (e ErrPropertyNames) Map(f func(Error) Error) Error {
	e.Err = f(e.Err)
	return e
}

// A constraint error returned by [MinProperties].
type ErrMinProperties struct {
	Format string
	Value  int
}

// GetDefault implements [Error] interface.
func (e ErrMinProperties) GetDefault() Error {
	return ErrMinProperties{}
}

// SetFormat implements [Error] interface.
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

// A constraint error returned by [MaxProperties].
type ErrMaxProperties struct {
	Format string
	Value  int
}

// GetDefault implements [Error] interface.
func (e ErrMaxProperties) GetDefault() Error {
	return ErrMaxProperties{}
}

// SetFormat implements [Error] interface.
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
