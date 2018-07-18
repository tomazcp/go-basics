package main

import (
	"fmt"
)

func main() {
	x := 5
	// we're not passing the memory address
	// we're passing the value
	zero(x)
	fmt.Println(x) // x is still 5
}

func zero(x int) {
	x = 0
}
