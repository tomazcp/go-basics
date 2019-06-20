package main

import (
	"fmt"
)

// untyped constant
const outer = "outer string"

// typed constant
const str string = "my string"

// multi declaration
const (
	a = 42
	b = 42.78
	c = "James Bond"
) // untyped constants are called constants of a kind (gives flexibility to the compiler)

func main() {
	const inner = "inner string"

	fmt.Println("outer - ", outer)
	fmt.Println("inner - ", inner)
}
