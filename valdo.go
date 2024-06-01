package valdo

func F[T any](field T, checks ...FieldCheck[T]) error {
	panic("todo")
}

type FieldChecks[T any] struct {
	checks []FieldCheck[T]
}

type FieldCheck[T any] func(T) *FieldError

func NewFieldChecks[T any](checks ...FieldCheck[T]) FieldChecks[T] {
	return FieldChecks[T]{checks}
}

func (cs FieldChecks[T]) Check(v any) Errors {
	panic("todo")
}
