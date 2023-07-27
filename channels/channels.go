package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var (
		stack      = make([]*tree.Tree, 100)
		start, end = 0, 0
	)
	done := false
	for !done {
		if t != nil {
			// push to stack
			stack[end] = t
			t = t.Left
			end++
		} else {
			// empty stack
			if start == end {
				done = true
			} else {
				pop := stack[end-1]
				end--
				ch <- pop.Value
				t = pop.Right
			}
		}
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1, v2 := <-ch1, <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		var wg sync.WaitGroup
		go func() {
			wg.Add(1)
			defer wg.Done()
			for in1 := range input1 {
				c <- in1
			}
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			for in2 := range input2 {
				c <- in2
			}
		}()
		wg.Wait()
		close(c)
	}()
	return c
}

type Ball struct{ Hits int }

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})
	ch5 := make(chan struct{})
	go func() {
		defer close(ch5)
	}()
	<-orChannel(ch1, ch2, ch3, ch4, ch5)
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.Hits++
		fmt.Println(name, ball.Hits)
		time.Sleep(time.Millisecond * 100)
		table <- ball
	}
}

func fib(n int) int {
	switch n {
	case 1, 2:
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func orChannel(receives ...<-chan struct{}) <-chan struct{} {
	switch len(receives) {
	case 0:
		return nil
	case 1:
		return receives[0]
	}

	orDone := make(chan struct{})
	go func() {
		defer close(orDone)

		switch len(receives) {
		case 2:
			select {
			case <-receives[0]:
			case <-receives[1]:
			}
		default:
			select {
			case <-receives[0]:
			case <-receives[1]:
			case <-receives[2]:
			case <-orChannel(append(receives[3:], orDone)...):
			}
		}
	}()
	fmt.Println("call orChannel")
	return orDone
}
