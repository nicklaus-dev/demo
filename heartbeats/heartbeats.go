package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		wg     sync.WaitGroup
		done   = make(chan struct{})
		result = make(chan int)
	)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go replicateRequests(done, i, &wg, result)
	}

	// read the first result
	firstResult := <-result
	close(done)
	wg.Wait()
	fmt.Printf("Received an answer from #%v\n", firstResult)
}

func doWork(done <-chan any, pulseInterval time.Duration) (<-chan any, <-chan time.Time) {
	heartbeats := make(chan any)
	results := make(chan time.Time)
	go func() {
		// defer func() {
		// 	close(heartbeats)
		// 	close(results)
		// }()
		pulse := time.Tick(pulseInterval)
	
		workGen := time.Tick(2 * pulseInterval)

		sendPulse := func() {
			select {
			case heartbeats <- struct{}{}:
			default:
			}
		}

		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		for i := 0; i < 2; i++ {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-workGen:
				sendResult(r)
			}
		}
	}()

	return heartbeats, results
}

func replicateRequests(done <-chan struct{}, id int, wg *sync.WaitGroup, result chan<- int) {
	start := time.Now()
	defer wg.Done()

	simulateLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
	select {
	case <-done:
	case <-time.After(simulateLoadTime):
	}

	select {
	case <-done:
	case result <- id:
	}
 
	took := time.Since(start)
	if took < simulateLoadTime {
		took = simulateLoadTime
	}
	
	fmt.Printf("%v took %v\n", id, took)
}
