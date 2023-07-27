package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int

	increment := func() { counter++ }
	decrement := func() { counter-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Println(counter)
}
