package main

import (
	"fmt"
)

// untyped constant
const outer = "outer string"

// typed constant
const str string = "my string"

func main() {
	const inner = "inner string"

	fmt.Println("outer - ", outer)
	fmt.Println("inner - ", inner)
}
