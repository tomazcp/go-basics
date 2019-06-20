package main

import "fmt"

const (
	a = iota
	b
	c
)

func main() {
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
}
