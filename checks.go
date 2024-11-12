package valdo

import (
	"bytes"
	"fmt"
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
	return FieldCheck[T]{
		check:   c,
		message: "must be greater than or equal to %v",
	}
}

func LTE[T internal.Ordered](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f <= v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be less than or equal to %v",
	}
}

func GT[T internal.Ordered](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f > v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be greater than %v",
	}
}

func LT[T internal.Ordered](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f < v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be less than %v",
	}
}

func Eq[T comparable](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f == v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be equal to %v",
	}
}

func NE[T comparable](v T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f != v {
			return nil
		}
		return newFieldError(m, v)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must not be equal to %v",
	}
}

func Positive[T internal.Integer | internal.Float]() FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f > 0 {
			return nil
		}
		return newFieldError(m)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be positive",
	}
}

func Matches[T ~string](p string) FieldCheck[T] {
	rex := regexp.MustCompile(p)
	c := func(m string, f T) *FieldError {
		if rex.MatchString(string(f)) {
			return nil
		}
		return newFieldError(m, p)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must match regular expression %v",
	}
}

func MatchesB[T ~[]byte](p string) FieldCheck[T] {
	rex := regexp.MustCompile(p)
	c := func(m string, f T) *FieldError {
		if rex.Match([]byte(f)) {
			return nil
		}
		return newFieldError(m, p)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must match regular expression %v",
	}
}

func Contains[T ~string](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if strings.Contains(string(f), string(sub)) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must contain %v",
	}
}

func ContainsB[T ~[]byte](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if bytes.Contains(f, sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must contain %v",
	}
}

func ContainsA[T ~[]E, E comparable](sub E) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if slices.Contains(f, sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must contain %v",
	}
}

func HasPrefix[T ~string](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if strings.HasPrefix(string(f), string(sub)) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must start with %v",
	}
}

func HasPrefixB[T ~[]byte](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if bytes.HasPrefix(f, sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must start with %v",
	}
}

func HasSuffix[T ~string](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if strings.HasSuffix(string(f), string(sub)) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must end with %v",
	}
}

func HasSuffixB[T ~[]byte](sub T) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if bytes.HasSuffix(f, sub) {
			return nil
		}
		return newFieldError(m, sub)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must end with %v",
	}
}

func MinLen[T ~string](min int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) >= min {
			return nil
		}
		return newFieldError(m, min)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be at least %d character(s)",
	}
}

func MinLenA[T ~[]E, E any](min int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) >= min {
			return nil
		}
		return newFieldError(m, min)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must have at least %d item(s)",
	}
}

func MaxLen[T ~string, E any](max int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) <= max {
			return nil
		}
		return newFieldError(m, max)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be at most %d character(s)",
	}
}

func MaxLenA[T ~[]E, E any](max int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) <= max {
			return nil
		}
		return newFieldError(m, max)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must have at most %d item(s)",
	}
}

func LenIs[T ~string](l int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) == l {
			return nil
		}
		return newFieldError(m, l)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be %d character(s)",
	}
}

func LenIsA[T ~[]E, E any](l int) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if len(f) == l {
			return nil
		}
		return newFieldError(m, l)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must have %d item(s)",
	}
}

func NotDefault[T comparable]() FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if f == *new(T) {
			return nil
		}
		return newFieldError(m)
	}
	return FieldCheck[T]{
		check:   c,
		message: "required",
	}
}

func Sorted[T ~[]E, E internal.Ordered]() FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if slices.IsSorted(f) {
			return nil
		}
		return newFieldError(m)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be sorted",
	}
}

func True[T any](fn func(T) bool) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if fn(f) {
			return nil
		}
		return newFieldError(m)
	}
	return FieldCheck[T]{
		check:   c,
		message: "invalid",
	}
}

func False[T any](fn func(T) bool) FieldCheck[T] {
	c := func(m string, f T) *FieldError {
		if !fn(f) {
			return nil
		}
		return newFieldError(m)
	}
	return FieldCheck[T]{
		check:   c,
		message: "invalid",
	}
}

func OneOf[T comparable](items ...T) FieldCheck[T] {
	stringed := make([]string, 0, len(items))
	for _, i := range items {
		stringed = append(stringed, fmt.Sprintf("%v", i))
	}
	joined := strings.Join(stringed, ", ")

	c := func(m string, f T) *FieldError {
		for _, i := range items {
			if i == f {
				return nil
			}
		}
		return newFieldError(m, joined)
	}
	return FieldCheck[T]{
		check:   c,
		message: "must be one of: %v",
	}
}
