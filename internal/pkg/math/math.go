package math

import (
	"errors"

	"github.com/felix-kaestner/calculator/internal/pkg/proto"
)

// ErrDivisionByZero means that a divisor of zero was
// encountered in the process of solving a mathematical equation.
var ErrDivisionByZero = errors.New("division by zero")

// ErrInvalidMethod means that the povided mathematical method
// operator is unsupported or invalid.
var ErrInvalidMethod = errors.New("invalid method")

// Implement the Operator interface for different mathemtical methods.
// Solve computes the result of applying the mathematical method m
// to the two operands a and b returning the result or an error is the
// operation is invalid.
func Solve(m proto.Method, a float64, b float64) (float64, error) {
	switch m {
	case proto.Method_ADD:
		return a + b, nil
	case proto.Method_SUB:
		return a - b, nil
	case proto.Method_MUL:
		return a * b, nil
	case proto.Method_DIV:
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, ErrInvalidMethod
	}
}
