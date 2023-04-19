package benchmark

import "testing"

func BenchMarkFoo(b *testing.B){
	for i := 0; i < b.N; i++ {
		foo()
	}
}

func foo() string {
	return "foo"
}