package assert

import (
	"testing"
)

func TestAssertNil(t *testing.T) {
	Nil(t, nil)

	i := (*interface{})(nil)
	Nil(t, i)

	var s []int
	Nil(t, s)

	var m map[int]int
	Nil(t, m)

	var f func()
	Nil(t, f)
}

func TestAssertEqualNil(t *testing.T) {
	Equal(t, nil, nil)

	i := (*interface{})(nil)
	Equal(t, i, i)

	var s []int
	Equal(t, s, s)

	var m map[int]int
	Equal(t, m, m)

	var f func()
	Equal(t, f, f)
}

func TestAssertEqual(t *testing.T) {
	Equal(t, "a", "a")
	Equal(t, 1, 1)
	Equal(t, []int{0}, []int{0})
	Equal(t, map[int]int{0: 1}, map[int]int{0: 1})
}
