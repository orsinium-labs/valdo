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

func Contains[T ~string](sub string) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if strings.Contains(string(f), sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{check: c, message: "must contain %v"}
}

func HasPrefix[T ~string](sub string) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if strings.HasPrefix(string(f), sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{check: c, message: "must start with %v"}
}

func HasSuffix[T ~string](sub string) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if strings.HasSuffix(string(f), sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{check: c, message: "must end with %v"}
}

func MinLen[T ~string](min int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) >= min {
			return nil
		}
		return newFieldError(m, min)
	}
	return FieldCheck[T]{check: c, message: "must be at least %d character(s)"}
}

func MaxLen[T ~string | ~[]E, E any](max int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) <= max {
			return nil
		}
		return newFieldError(m, max)
	}
	return FieldCheck[T]{check: c, message: "must be at most %d character(s)"}
}

func LenIs[T ~string | ~[]E, E any](l int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) == l {
			return nil
		}
		return newFieldError(m, l)
	}
	return FieldCheck[T]{check: c, message: "must be %d character(s)"}
}

func NotDefault[T comparable]() FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f == *new(T) {
			return nil
		}
		return newFieldError(m)
	}
	return FieldCheck[T]{check: c, message: "must not be default value"}
}
