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

func Eq[T comparable](v T) FieldCheck[T] {
	return func(f T) *FieldError {
		if f == v {
			return nil
		}
		return newFieldError("must be equal to %v", v)
	}
}

func Positive[T internal.Integer | internal.Float](f T) *FieldError {
	if f > 0 {
		return nil
	}
	return newFieldError("must be positive")
}

func Matches[T ~string](p string) FieldCheck[T] {
	rex := regexp.MustCompile(p)
	return func(f T) *FieldError {
		if rex.MatchString(string(f)) {
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
