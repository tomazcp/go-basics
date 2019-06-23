package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)  // 42
	fmt.Println(&a) // address

	b := &a
	fmt.Println(b)   // address
	fmt.Println(*b)  // dereference 42 -> * gives you the value stored
	fmt.Println(*&a) // deference 42

	test(&a)
	fmt.Println(a)  // 55
	fmt.Println(*b) // 55 -> b has the same address as a
}

// changing the value at the address
func test(a *int) {
	*a = 55
}
