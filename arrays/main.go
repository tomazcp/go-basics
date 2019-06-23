package main

import "fmt"

// don't use arrays, use slices
func main() {
	var x [5]int   // the array length is part of its type
	fmt.Println(x) // [0 0 0 0 0] zero value
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x)) // length
	fmt.Printf("%T", x)
}
