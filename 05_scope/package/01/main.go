package main

import "fmt"

// package level scope meaning it would be visible to other
// files in the same package
var x = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
