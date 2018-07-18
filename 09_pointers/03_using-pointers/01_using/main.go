package main

import (
	"fmt"
)

func zero(z *int) {
	*z = 0 // dereferencing
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0 because we dereferenced it inside of func zero()
}
