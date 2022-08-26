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
