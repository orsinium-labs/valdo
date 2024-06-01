package valdo

import (
	"regexp"
	"strings"

	"github.com/orsinium-labs/valdo/internal"
)

func GTE[T internal.Ordered](v T) FieldCheck[T] {
	return func(f T) *FieldError {
		if f >= v {
			return nil
		}
		return newFieldError("must be greater than or equal to %v", v)
	}
}

func LTE[T internal.Ordered](v T) FieldCheck[T] {
	return func(f T) *FieldError {
		if f <= v {
			return nil
		}
		return newFieldError("must be less than or equal to %v", v)
	}
}

func GT[T internal.Ordered](v T) FieldCheck[T] {
	return func(f T) *FieldError {
		if f > v {
			return nil
		}
		return newFieldError("must be greater than %v", v)
	}
}

func LT[T internal.Ordered](v T) FieldCheck[T] {
	return func(f T) *FieldError {
		if f < v {
			return nil
		}
		return newFieldError("must be less than %v", v)
	}
}

func Eq[T internal.Ordered](v T) FieldCheck[T] {
	return func(f T) *FieldError {
		if f == v {
			return nil
		}
		return newFieldError("must be equal to %v", v)
	}
}

func Matches(p string) FieldCheck[string] {
	rex := regexp.MustCompile(p)
	return func(f string) *FieldError {
		if rex.MatchString(f) {
			return nil
		}
		return newFieldError("must match regular expression %v", p)
	}
}

func Contains(sub string) FieldCheck[string] {
	return func(f string) *FieldError {
		if strings.Contains(f, sub) {
			return nil
		}
		return newFieldError("must contain %v", sub)
	}
}

func MinLen(m int) FieldCheck[string] {
	return func(f string) *FieldError {
		if len(f) >= m {
			return nil
		}
		return newFieldError("must be at least %d character(s)", m)
	}
}

func MaxLen(m int) FieldCheck[string] {
	return func(f string) *FieldError {
		if len(f) <= m {
			return nil
		}
		return newFieldError("must be at most %d character(s)", m)
	}
}
