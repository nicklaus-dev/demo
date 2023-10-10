package main

import "fmt"

func main() {
	m := map[string]int{"one": 1}
	fmt.Printf("before:%v\n", m)
	clear(m)
	fmt.Printf("after:%v\n", m)
}
