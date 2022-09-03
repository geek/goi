package goi

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

type str struct {
	*Any[string]
	min   *uint64
	max   *uint64
	regex *regexp.Regexp
}

func String(name string) str {
	return str{Any: &Any[string]{name: name}}
}

// Marks the field as requiring a empty string value
func (s str) Required() str {
	a := s.Any.Required()
	s.Any = &a
	return s
}

// Valid only allows values that match the provided valids args
func (s str) Valid(valids ...string) str {
	a := s.Any.Valid(valids...)
	s.Any = &a
	return s
}

// Invalid only prevents values that match the provided invalids args
func (s str) Invalid(invalids ...string) str {
	a := s.Any.Invalid(invalids...)
	s.Any = &a
	return s
}

// Min specifies the minimum length allowed
func (s str) Min(m uint64) str {
	s.min = &m
	return s
}

// Max specifies the maximum length allowed
func (s str) Max(m uint64) str {
	if s.min != nil && m < *s.min {
		panic("max cannot be less than min")
	}
	s.max = &m
	return s
}

func (s str) Regex(regex string) str {
	s.regex = regexp.MustCompile(regex)
	return s
}

func (s str) Validate(input string) error {
	err := s.Any.Validate(input)
	if err != nil {
		return err
	}

	l := utf8.RuneCountInString(input)
	if s.max != nil && l > int(*s.max) {
		return fmt.Errorf("%s cannot be greater in length than %d", s.name, s.max)
	}

	if s.min != nil && l < int(*s.min) {
		return NewValidationError(s.name, "%s should be at least %d in length", s.name, s.min)
	}

	b := []byte(input)
	if s.regex != nil && !s.regex.Match(b) {
		return NewValidationError(s.name, "%s should match expression %s", s.name, s.regex.String())
	}

	return nil
}
