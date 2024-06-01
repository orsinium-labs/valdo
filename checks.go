package valdo

import (
	"bytes"
	"regexp"
	"slices"
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

func NE[T comparable](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f != v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{check: c, message: "must not be equal to %v"}
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

func BMatches[T ~[]byte](p string) FieldCheck[T] {
	rex := regexp.MustCompile(p)
	c := func(m string, f T) *FieldError {
		if rex.Match([]byte(f)) {
			return nil
		}
		return newFieldError(m, p)
	}
	return FieldCheck[T]{check: c, message: "must match regular expression %v"}
}

func Contains[T ~string](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if strings.Contains(string(f), string(sub)) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{check: c, message: "must contain %v"}
}

func BContains[T ~[]byte](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if bytes.Contains(f, sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{check: c, message: "must contain %v"}
}

func SContains[T ~[]E, E comparable](sub E) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if slices.Contains(f, sub) {
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

func BHasPrefix[T ~[]byte](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if bytes.HasPrefix(f, sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{check: c, message: "must start with %v"}
}

func HasSuffix[T ~string](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if strings.HasSuffix(string(f), string(sub)) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{check: c, message: "must end with %v"}
}

func BHasSuffix[T ~[]byte](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if bytes.HasSuffix(f, sub) {
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

func SMinLen[T ~[]E, E any](min int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) >= min {
			return nil
		}
		return newFieldError(m, min)
	}
	return FieldCheck[T]{check: c, message: "must have at least %d item(s)"}
}

func MaxLen[T ~string, E any](max int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) <= max {
			return nil
		}
		return newFieldError(m, max)
	}
	return FieldCheck[T]{check: c, message: "must be at most %d character(s)"}
}

func SMaxLen[T ~[]E, E any](max int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) <= max {
			return nil
		}
		return newFieldError(m, max)
	}
	return FieldCheck[T]{check: c, message: "must have at most %d item(s)"}
}

func LenIs[T ~string](l int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) == l {
			return nil
		}
		return newFieldError(m, l)
	}
	return FieldCheck[T]{check: c, message: "must be %d character(s)"}
}

func SLenIs[T ~[]E, E any](l int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) == l {
			return nil
		}
		return newFieldError(m, l)
	}
	return FieldCheck[T]{check: c, message: "must have %d item(s)"}
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

func Sorted[T ~[]E, E internal.Ordered]() FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if slices.IsSorted(f) {
			return nil
		}
		return newFieldError(m)
	}
	return FieldCheck[T]{check: c, message: "must be sorted"}
}
