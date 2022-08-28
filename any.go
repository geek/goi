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
	valids     []T
	invalids   []T
}

func New[T comparable](name string) Any[T] {
	return Any[T]{name: name}
}

func (a Any[T]) Required() Any[T] {
	a.isRequired = true
	return a
}

func (a Any[T]) Valid(valids ...T) Any[T] {
	a.valids = valids
	return a
}

func (a Any[T]) Invalid(invalids ...T) Any[T] {
	a.invalids = invalids
	return a
}

func (a Any[T]) Validate(input T) error {
	if a.isRequired && isZero(input) {
		return NewValidationError(a.name, "%s is required", a.name)
	}

	if len(a.valids) > 0 {
		isValid := false
		for _, valid := range a.valids {
			if input == valid {
				isValid = true
				break
			}
		}

		if !isValid {
			return NewValidationError(a.name, "must be one of the following values: %v", a.valids)
		}
	}

	if len(a.invalids) > 0 {
		isValid := true
		for _, invalid := range a.invalids {
			if input == invalid {
				isValid = false
				break
			}
		}

		if !isValid {
			return NewValidationError(a.name, "must not equal any of the following values: %v", a.valids)
		}
	}

	return nil
}

func isZero[T comparable](v T) bool {
	return v == *new(T)
}
