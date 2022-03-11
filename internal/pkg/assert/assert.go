package assert

import (
	"reflect"
	"testing"
)

// isNil checks if a value is nil.
// This method accounts for nil pointers and nil interfaces.
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	}

	return false
}

// Nil asserts that a value is nil.
func Nil(t *testing.T, actual interface{}) {
	if !isNil(actual) {
		t.Errorf("Test %s: Expected value to be nil, Received `%v` (type %v)", t.Name(), actual, reflect.TypeOf(actual))
	}
}

// Equal asserts that two values are equal.
func Equal(t *testing.T, expected, actual interface{}) {
	if (isNil(expected) && isNil(actual)) || reflect.DeepEqual(expected, actual) {
		return
	}

	t.Errorf("Test %s: Expected `%v` (type %v), Received `%v` (type %v)", t.Name(), expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
}
