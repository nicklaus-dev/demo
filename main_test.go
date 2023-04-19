package main

import (
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func sum8() int64 {
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	// if cpu cache is 64 bytes
	var sum int64
	for i := 0; i < len(nums); i += 8 {
		sum += nums[i]
	}

	return sum
}

type Input struct {
	a int64
	b int64
}

type Result struct {
	sumA int64
	sumB int64
}

func count(inputs []Input) Result {
	var wg sync.WaitGroup
	wg.Add(2)
	result := Result{}
	go func() {
		defer wg.Done()
		for _, it := range inputs {
			result.sumA += it.a
		}
	}()

	go func() {
		defer wg.Done()
		for _, it := range inputs {
			result.sumB += it.b
		}
	}()

	wg.Wait()
	return result
}
