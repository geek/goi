package goi

import "fmt"

type number interface {
	int64 | int32 | uint64 | uint32
}

type num[T number] struct {
	*Any[T]
	min *T
	max *T
}

func (n num[T]) Required() num[T] {
	n.isRequired = true
	return n
}

func Number[T number](name string) num[T] {
	return num[T]{Any: &Any[T]{name, false}}
}

func (n num[T]) Min(m T) num[T] {
	n.min = &m
	return n
}

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
