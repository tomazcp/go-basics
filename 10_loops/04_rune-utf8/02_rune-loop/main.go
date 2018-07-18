package main

import "fmt"

func main() {
	foo := 'a'

	fmt.Println(foo)         // prints 97
	fmt.Printf("%T \n", foo) // %T shows the type (int32)
	// rune is an alias for int32
	// print int32
	fmt.Printf("%T \n", rune(foo))
}
