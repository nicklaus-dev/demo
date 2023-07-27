package main

func Hello(s string) string {
	panic("hello, there is a panic")
	return "Hello, " + s
}