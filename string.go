package goi

import (
	"fmt"
	"regexp"
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

func (s str) Required() str {
	s.isRequired = true
	return s
}

func (s str) Valid(valids ...string) str {
	s.valids = valids
	return s
}

func (s str) Invalid(invalids ...string) str {
	s.invalids = invalids
	return s
}

func (s str) Min(m uint64) str {
	s.min = &m
	return s
}

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

	if s.max != nil && len(input) > int(*s.max) {
		return fmt.Errorf("%s cannot be greater in length than %d", s.name, s.max)
	}

	if s.min != nil && len(input) < int(*s.min) {
		return NewValidationError(s.name, "%s should be at least %d in length", s.name, s.min)
	}

	b := []byte(input)
	if s.regex != nil && !s.regex.Match(b) {
		return NewValidationError(s.name, "%s should match expression %s", s.name, s.regex.String())
	}

	return nil
}
