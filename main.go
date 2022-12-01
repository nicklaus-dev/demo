package main

import (
	"fmt"
)

func main() {
	i := 0
	c := make(chan struct{})
	go func() {
		fmt.Printf("i: %v\n", i)
		<-c
	}()
	i++
	close(c)
}
