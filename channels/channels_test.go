package main

import "testing"

func TestSlice(t *testing.T) {
	slices := []int{1, 2, 3}

	t.Logf("%v", slices[3:])
}
