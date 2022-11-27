package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var i int64
	var mx sync.Mutex
	go func() {
		mx.Lock()
		defer mx.Unlock()
		i++
	}()

	go func() {
		mx.Lock()
		defer mx.Unlock()
		i++
	}()

	time.Sleep(time.Second)
	fmt.Printf("i: %v\n", i)
}
