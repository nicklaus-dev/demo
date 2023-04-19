package main

import (
	"io"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func main() {

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

func BenchMarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
}

func foo() string {
	return "foo"
}