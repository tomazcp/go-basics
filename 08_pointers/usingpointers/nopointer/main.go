package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a - ", &a)

	// b value type is of int pointer
	var b *int = &a

	// this will print the memory address to which b is pointing to
	fmt.Println(b)

	// here we are dereferencing b so now b will print out the value 42
	fmt.Println(*b)

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42

	// this is useful
	// we can pass a memory address instead of a bunch of values (we pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// in Go everything is PASS BY VALUE
	// so when we are passing a memory address we are passing a value
}
