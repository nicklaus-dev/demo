package main

import "fmt"

func main() {
	arr := []int{1, 2}
	fmt.Printf("%v\n", arr[:len(arr)-1])
}
