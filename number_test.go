package goi_test

import (
	"testing"

	"github.com/geek/goi"
	"github.com/stretchr/testify/assert"
)

func TestNumber_MinMaxUInt(t *testing.T) {
	n := goi.Number[uint64]("test").Required().Min(5).Max(10)

	assert.NoError(t, n.Validate(5))
	assert.Error(t, n.Validate(0))
	assert.Error(t, n.Validate(1))
	assert.Error(t, n.Validate(11))
}

func TestNumber_MinMaxInt(t *testing.T) {
	n := goi.Number[int64]("test").Required().Min(-5).Max(10)

	assert.NoError(t, n.Validate(5))
	assert.Error(t, n.Validate(0))
	assert.Error(t, n.Validate(-6))
	assert.Error(t, n.Validate(11))
}
