package main

import "fmt"

func main() {
	a := 42

	fmt.Println("a - ", a)
	// address is hexadecimal
	fmt.Println("a - ", &a)
	// converting address to decimal
	fmt.Printf("%d \n", &a)
}
