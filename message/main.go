package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		meesageCh    = make(chan struct{})
		disconnectCh = make(chan struct{})
	)
	defer func() {
		close(meesageCh)
		close(disconnectCh)
	}()

	go func() {
		for {
			select {
			case v := <-meesageCh:
				fmt.Printf("v: %v\n", v)
			case <-disconnectCh:
				fmt.Println("disconnection, return")
				return
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			meesageCh <- struct{}{}
		}
		disconnectCh <- struct{}{}
	}()

	time.Sleep(time.Second * 5)
}
