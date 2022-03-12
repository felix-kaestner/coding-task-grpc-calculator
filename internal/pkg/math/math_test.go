package math

import (
	"errors"
	"reflect"
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

func TestDivisionByZero(t *testing.T) {
	_, err := Solve(proto.Method_DIV, 3.0, 0.0)
	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("Test %s: Expected `%v` (type %v), Received `%v` (type %v)", t.Name(), ErrDivisionByZero, reflect.TypeOf(ErrDivisionByZero), err, reflect.TypeOf(err))
	}
}

func TestInvalidMethod(t *testing.T) {
	_, err := Solve(4, 1.0, 2.0)
	if !errors.Is(err, ErrInvalidMethod) {
		t.Errorf("Test %s: Expected `%v` (type %v), Received `%v` (type %v)", t.Name(), ErrInvalidMethod, reflect.TypeOf(ErrInvalidMethod), err, reflect.TypeOf(err))
	}
}
