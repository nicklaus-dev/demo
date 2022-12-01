package concurrent

import (
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	var i int
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch
	i += <-ch

	time.Sleep(time.Second)
	t.Log(i)
}

func TestRaceCondition(t *testing.T) {
	var i int
	var mx sync.Mutex
	go func() {
		mx.Lock()
		defer mx.Unlock()
		i = 1
	}()
	go func() {
		mx.Lock()
		defer mx.Unlock()
		i = 2
	}()
	time.Sleep(time.Second)
	t.Log(i)
}
