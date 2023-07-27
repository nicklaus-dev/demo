package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("Create a new instance.")
			return struct{}{}
		},
	}

	a := pool.Get()
	a2 := pool.Get()
	pool.Put(struct{}{})
	a3 := pool.Get()
	fmt.Println(a, a2, a3)
}
