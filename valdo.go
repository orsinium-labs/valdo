package valdo

type FieldChecks[T any] struct {
	checks []FieldCheck[T]
}

func F[T any](checks ...FieldCheck[T]) FieldChecks[T] {
	return FieldChecks[T]{checks}
}

func (cs FieldChecks[T]) Check(v T) Errors {
	panic("todo")
}

// Make a copy with the given additional checks.
func (cs FieldChecks[T]) With(checks ...FieldCheck[T]) FieldChecks[T] {
	result := append([]FieldCheck[T]{}, cs.checks...)
	result = append(result, checks...)
	cs.checks = result
	return cs
}

type FieldCheck[T any] struct {
	check   func(string, T) *FieldError
	message string
}

// Make a copy of the check with the given error message.
func (f FieldCheck[T]) M(message string) FieldCheck[T] {
	f.message = message
	return f
}
