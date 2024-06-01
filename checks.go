package valdo

import (
	"regexp"
	"strings"

	"github.com/orsinium-labs/valdo/internal"
)

func GTE[T internal.Ordered](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f >= v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{check: c, message: "must be greater than or equal to %v"}
}

func LTE[T internal.Ordered](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f <= v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{check: c, message: "must be less than or equal to %v"}
}

func GT[T internal.Ordered](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f > v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{check: c, message: "must be greater than %v"}
}

func LT[T internal.Ordered](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f < v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{check: c, message: "must be less than %v"}
}

func Eq[T comparable](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f == v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{check: c, message: "must be equal to %v"}
}

func Positive[T internal.Integer | internal.Float]() FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f > 0 {
			return nil
		}
		return newFieldError(m)
	}
	return FieldCheck[T]{check: c, message: "must be positive"}
}

func Matches[T ~string](p string) FieldCheck[T] {
	rex := regexp.MustCompile(p)
	c := func(m string, f T) *FieldError {
		if rex.MatchString(string(f)) {
			return nil
		}
		return newFieldError(m, p)
	}
	return FieldCheck[T]{check: c, message: "must match regular expression %v"}
}

func Contains(sub string) FieldCheck[string] {
	c := func(m string, f string) *FieldError {
		if strings.Contains(f, sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[string]{check: c, message: "must contain %v"}
}

func MinLen(min int) FieldCheck[string] {
	c := func(m string, f string) *FieldError {
		if len(f) >= min {
			return nil
		}
		return newFieldError(m, min)
	}
	return FieldCheck[string]{check: c, message: "must be at least %d character(s)"}
}

func MaxLen(max int) FieldCheck[string] {
	c := func(m string, f string) *FieldError {
		if len(f) <= max {
			return nil
		}
		return newFieldError(m, max)
	}
	return FieldCheck[string]{check: c, message: "must be at most %d character(s)"}
}

func LenIs(l int) FieldCheck[string] {
	c := func(m string, f string) *FieldError {
		if len(f) == l {
			return nil
		}
		return newFieldError(m, l)
	}
	return FieldCheck[string]{check: c, message: "must be %d character(s)"}
}
