package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{counters: map[string]int{}}
}

func (c *Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := NewCounter()

	go c.Increment("nick")
	go c.Increment("tom")

	time.Sleep(time.Second)
	fmt.Printf("c.counters: %v\n", c.counters)
}
