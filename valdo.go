package valdo

type FieldChecks[T any] struct {
	checks []FieldCheck[T]
}

func F[T any](checks ...FieldCheck[T]) FieldChecks[T] {
	return FieldChecks[T]{checks}
}

func (cs FieldChecks[T]) Check(v any) Errors {
	panic("todo")
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
