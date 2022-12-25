package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func(v int) {
			fmt.Printf("i: %v\n", v)
		}(i)
	}
	time.Sleep(time.Second)
}
