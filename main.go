package main

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	w := newWatcher()
	defer w.close()
	n := 1000
	i := runtime.NumCPU()
	fmt.Printf("i: %v\n", i)
	wg := sync.WaitGroup{}
	wg.Add(n)
	start := time.Now()
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			handle()
		}()
	}
	wg.Wait()
	end := time.Now()
	cost := end.Sub(start).Milliseconds()
	fmt.Printf("cost: %v\n", cost)
}

func handle() {
	time.Sleep(time.Second)
}

func read(r io.Reader) (int, error) {
	var count int64
	wg := sync.WaitGroup{}
	n := 10

	ch := make(chan []byte, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, v)
			}
		}()
	}

	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err == io.EOF {
			break
		}
		ch <- b
	}
	close(ch)
	wg.Wait()
	return int(count), nil
}

func task(b []byte) int64 {
	return 1
}

type watcher struct {
}

func (w watcher) watch() {
}

func (w watcher) close() {
	// close all resource
}

func newWatcher() watcher {
	w := watcher{}
	go w.watch()
	return w
}
