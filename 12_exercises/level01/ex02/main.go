package main

import "fmt"

// exercício 3 aqui dentro também
// declarado desta forma para fins de exercício
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// zero-values
	fmt.Println(x) // 0
	fmt.Println(y) // ""
	fmt.Println(z) // false
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
