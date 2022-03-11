package math

import (
	"testing"

	"github.com/felix-kaestner/calculator/internal/pkg/assert"
	"github.com/felix-kaestner/calculator/internal/pkg/proto"
)

func TestAddition(t *testing.T) {
	res, err := Solve(proto.Method_ADD, 1.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, res)
}

func TestSubtraction(t *testing.T) {
	res, err := Solve(proto.Method_SUB, 3.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, res)
}

func TestMultiplication(t *testing.T) {
	res, err := Solve(proto.Method_MUL, 2.0, 3.0)
	assert.Nil(t, err)
	assert.Equal(t, 6.0, res)
}

func TestDivision(t *testing.T) {
	res, err := Solve(proto.Method_DIV, 3.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, 1.5, res)
}
