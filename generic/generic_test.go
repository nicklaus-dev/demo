package generic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	t.Cleanup(func() {
		t.Log("TestAny finish!")
	})
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

func filter[T any](list []T, filterFn func(T) bool) []T {
	var result []T
	for _, v := range list {
		if filterFn(v) {
			result = append(result, v)
		}
	}
	return result
}

func makeGo() {	
	ch := make(chan int)

	for v := range ch {
		_ = v
	}

	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
	}
}
