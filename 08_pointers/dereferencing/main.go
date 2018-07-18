package main

import (
	"fmt"
)

func main() {
	a := 42

	fmt.Println("a - ", a)
	fmt.Println("a - ", &a)

	var b *int = &a // b value type is of int pointer

	fmt.Println(b) // this will print the memory address

	// here we are dereferencing b so now b will print out the value 42
	fmt.Println(*b)

}
