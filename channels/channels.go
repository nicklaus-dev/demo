package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s interface{}
	fmt.Printf("unsafe.Sizeof(s): %v\n", unsafe.Sizeof(s))
	var st struct{}
	fmt.Printf("unsafe.Sizeof(st): %v\n", unsafe.Sizeof(st))

	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	ch <- struct{}{}
}
