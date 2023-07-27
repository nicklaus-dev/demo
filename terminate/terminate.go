package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(strs <-chan string, done <-chan struct{}) <-chan struct{} {
		terminate := make(chan struct{})
		go func() {
			defer close(terminate)

			for {
				select {
				case s := <-strs:
					fmt.Println("Received:", s)
				case <-done:
					fmt.Println("DoWork is terminated!")
					return
				}
			}
		}()
		return terminate
	}

	done := make(chan struct{})

	teriminate := doWork(nil, done)

	go func() {
		defer close(done)
		time.Sleep(time.Second)
	}()

	<-teriminate
	fmt.Println("It's all done!")
}
