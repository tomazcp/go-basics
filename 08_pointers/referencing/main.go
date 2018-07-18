package main

import (
	"fmt"
)

func main() {
	a := 42

	fmt.Println("a value is ", a)
	fmt.Println("a memory address is ", &a)

	var b *int = &a // b value type is of int pointer

	fmt.Println("b points to", b)

}
