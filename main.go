package main
 
import (
	"fmt"
	"io"

	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type MySlice []string

func (s MySlice) String() string {
	return strings.Join(s, "+")
}

func main() {
	m := MySlice{"1", "2", "3"}
	s := clone(m)
	fmt.Println(s)
}

func clone[S ~[]E, E any](s S) S {
	return s
}

// 1, 1, 2, 3, 5, 8, 13
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

func fibonacciV2(n uint) uint {
	if n < 2 {
		return n
	}
	return fibonacciV2(n-1) + fibonacciV2(n-2)
}

func doRequest() bool {
	time.Sleep(time.Millisecond * 50)
	return true
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
