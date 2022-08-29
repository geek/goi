package goi_test

import (
	"testing"

	"github.com/geek/goi"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	s := goi.New[string]("test").Required()

	assert.NoError(t, s.Validate("foo"))
	assert.Error(t, s.Validate(""))
}

func TestAny_Required(t *testing.T) {
	s := goi.New[bool]("test").Required()

	assert.NoError(t, s.Validate(true))
	assert.Error(t, s.Validate(false))
}
