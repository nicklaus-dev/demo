package main

import "fmt"

func main() {
	i := 0
	//ch := make(chan struct{}, 1) buffered channel
	ch := make(chan struct{}) // unbufferd channel
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Printf("i: %v\n", i)
}
