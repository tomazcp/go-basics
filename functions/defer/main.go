package main

import "fmt"

func main() {
	defer foo() // foo defered
	bar()       // bar runs first, then foo
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
