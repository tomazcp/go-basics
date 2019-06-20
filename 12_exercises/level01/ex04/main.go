package main

import "fmt"

type mytype int

var x mytype

// ex5
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	// ex5
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
