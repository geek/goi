package goi

import "fmt"

type number interface {
	int64 | int32 | uint64 | uint32 | float32 | float64
}

type num[T number] struct {
	*Any[T]
	min *T
	max *T
}

func Number[T number](name string) num[T] {
	return num[T]{Any: &Any[T]{name: name}}
}

// Marks the field as requiring a non-zero value
func (n num[T]) Required() num[T] {
	a := n.Any.Required()
	n.Any = &a
	return n
}

// Valid only allows values that match the provided valids args
func (n num[T]) Valid(valids ...T) num[T] {
	a := n.Any.Valid(valids...)
	n.Any = &a
	return n
}

// Invalid only prevents values that match the provided invalids args
func (n num[T]) Invalid(invalids ...T) num[T] {
	a := n.Any.Invalid(invalids...)
	n.Any = &a
	return n
}

// Min specifies the minimum value allowed
func (n num[T]) Min(m T) num[T] {
	n.min = &m
	return n
}

// Max specifies the maximum value allowed
func (n num[T]) Max(m T) num[T] {
	if n.min != nil && m < *n.min {
		panic("max cannot be less than min")
	}
	n.max = &m
	return n
}

func (n num[T]) Validate(input T) error {
	err := n.Any.Validate(input)
	if err != nil {
		return err
	}

	if n.max != nil && input > *n.max {
		return fmt.Errorf("%s cannot be greater in length than %d", n.name, n.max)
	}

	if n.min != nil && input < *n.min {
		return NewValidationError(n.name, "%s should be at least %d in length", n.name, n.min)
	}

	return nil
}
