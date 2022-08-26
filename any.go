package goi

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func NewValidationError(field, format string, a ...any) ValidationError {
	return ValidationError{
		Field:   field,
		Message: fmt.Sprintf(format, a...),
	}
}

func (v ValidationError) Error() string {
	return v.Message
}

type Any[T comparable] struct {
	name       string
	isRequired bool
}

func New[T comparable](name string) Any[T] {
	return Any[T]{name, false}
}

func (a Any[T]) Required() Any[T] {
	a.isRequired = true
	return a
}

func (a Any[T]) Validate(input T) error {
	if a.isRequired && isZero(input) {
		return NewValidationError(a.name, "%s is required", a.name)
	}

	return nil
}

func isZero[T comparable](v T) bool {
	return v == *new(T)
}
