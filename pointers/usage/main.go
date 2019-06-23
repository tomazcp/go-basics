package main

import "fmt"

func main() {
	x := 0
	foo(x)
	fmt.Println(x) // 0

	y := 0
	bar(&y)
	fmt.Println(y) // 43 changed from 0
}

func foo(n int) {
	fmt.Println(n) // 0
	n = 43
	fmt.Println(n) // 43
}

func bar(n *int) {
	fmt.Println(*n) // 0
	*n = 43
	fmt.Println(*n) // 43
}
