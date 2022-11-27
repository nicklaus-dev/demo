package concurrent

import (
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
