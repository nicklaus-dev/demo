package generic

import (
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	i := getAny("nick")
	itype := reflect.TypeOf(i)
	t.Log(itype)

	s, i2 := mapAny("strings", 1)
	t.Log(s)
	t.Log(i2)
}

func getAny[T any](value T) T {
	return value
}

func mapAny[K, T any](first K, second T) (K, T) {
	return first, second
}
