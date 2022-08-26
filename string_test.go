package goi_test

import (
	"testing"

	"github.com/geek/goi"
	"github.com/stretchr/testify/assert"
)

func TestString_MinMax(t *testing.T) {
	s := goi.String("test").Required().Min(1).Max(4)

	assert.NoError(t, s.Validate("foo"))
	assert.Error(t, s.Validate(""))
	assert.Error(t, s.Validate("12345"))
}

func TestString_Regex(t *testing.T) {
	s := goi.String("test").Required().Regex("^foo$")

	assert.NoError(t, s.Validate("foo"))
	assert.Error(t, s.Validate(""))
	assert.Error(t, s.Validate("12345"))
}
