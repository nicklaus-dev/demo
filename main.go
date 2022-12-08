package main

import (
	"fmt"
)

type Error interface{
	Msg() string
}

type User struct {
	Name string
	Age int
}

func (u User) Msg() string {
	return "invalid user name"
}

type Car struct {
	Brand string
}


func (c Car) Msg() string {
	return "invalid brand"
}

func checkMsg[E Error](e E) E {
	if e.Msg() == "invalid brand" {
		return e
	}

	// return empty e
	return *new(E)
}

func main() {
	u := User{
		Name: "nick",
	}
	c := Car{
		Brand: "benz",
	}
	fmt.Printf("before u: %v\n", u) // before u: {nick 0}
	fmt.Printf("before c: %v\n", c) // before c: {benz}
	u = checkMsg(u)
	c = checkMsg(c)
	fmt.Printf("after u: %v\n", u) // after u: { 0}
	fmt.Printf("after c: %v\n", c) // after c: {benz}
}
