package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var (
		memConsumed = func() uint64 {
			runtime.GC()
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			return memStats.Sys
		}

		c <-chan struct{}

		wg sync.WaitGroup

		noop = func() {
			wg.Done()
			<-c
		}

		numOfGorutinues = 10000
	)

	wg.Add(numOfGorutinues)
	before := memConsumed()
	for i := 0; i < numOfGorutinues; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("single gorutinue memory: %.3f kb\n", float64(after-before)/float64(numOfGorutinues)/1024)
}
