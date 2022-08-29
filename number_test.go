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

func TestNumber_MinMaxFloat(t *testing.T) {
	n := goi.Number[float64]("test").Required().Min(-.5).Max(1.5)

	assert.NoError(t, n.Validate(1.2))
	assert.Error(t, n.Validate(-.6))
	assert.Error(t, n.Validate(1.6))
}

func TestNumber_ValidFloat(t *testing.T) {
	n := goi.Number[float64]("test").Invalid(1.1, 1.6, -.6)

	assert.NoError(t, n.Validate(1.2))
	assert.Error(t, n.Validate(-.6))
	assert.Error(t, n.Validate(1.6))
}

func TestNumber_InvalidFloat(t *testing.T) {
	n := goi.Number[float64]("test").Invalid(1.1, 1.2)

	assert.NoError(t, n.Validate(1.0))
	assert.Error(t, n.Validate(1.1))
	assert.Error(t, n.Validate(1.2))
}

func TestNumber_Float(t *testing.T) {
	n := goi.Number[float64]("test").Invalid(1.1, 1.2).Min(.5).Max(3.0)

	assert.NoError(t, n.Validate(1.0))
	assert.NoError(t, n.Validate(2.2))
	assert.Error(t, n.Validate(1.1))
	assert.Error(t, n.Validate(1.2))
}
