package main

import (
	"fmt"
)

func main() {
	outerVal := outer()
	fmt.Println(outerVal())
}

// func() int is the return because we want to return
// the outer function to be able to assign it to a variable
func outer() func() int {
	a := 2
	return func() int {
		return a
	}
}
