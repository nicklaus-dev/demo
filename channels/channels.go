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
	// done := make(chan struct{})
	// defer close(done)

	// out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4))
	// for e := range out1 {
	// 	fmt.Printf("out1:%v out2:%v \n", e, <-out2)
	// }

	// done := make(chan struct{})
	// defer close(done)

	// ch1 := take(done, repeat(done, 1, 2), 4)

	// ch2 := make(chan any)
	// go func ()  {
	// 	for e := range ch1 {
	// 		var ch2 = ch2
	// 		ch2 <- e
	// 	}
	// }()

	// for e := range ch2 {
	// 	fmt.Println(e)
	// }

	ch := make(chan any)
	go func ()  {
		defer close(ch)
		ch <- 1
		var ch = ch
		ch <- 2
	}()

	for e := range ch {
		fmt.Println(e)
	}

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

func repeat(done <-chan struct{}, values ...any) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case ch <- v:

				}
			}
		}
	}()
	return ch
}

func take(done <-chan struct{}, inputStream <-chan any, num int) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case ch <- <-inputStream:

			}
		}
	}()
	return ch
}

func orDone(done <-chan struct{}, valStream <-chan any) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case val, ok := <-valStream:
				if !ok {
					return
				}
				select {
				case ch <- val:
				case <-done:
					return
				}
			}
		}
	}()
	return ch
}

func tee(done <-chan struct{}, ch <-chan any) (<-chan any, <-chan any) {
	out1, out2 := make(chan any), make(chan any)
	go func() {
		defer close(out1)
		defer close(out2)
		fmt.Println("before", out1, out2)
		for v := range orDone(done, ch) {
			var out1, out2 = out1, out2
			fmt.Println(out1)
			fmt.Println(out2)
			for i := 0; i < 2; i++ {
				select {
				case out1 <- v:
					out1 = nil
				case out2 <- v:
					out2 = nil
				}
			}
		}
		fmt.Println("after", out1, out2)
	}()
	return out1, out2
}
