package main

import "fmt"

func main() {
	current := incrementor()
	for i := 0; i < 50; i++ {
		fmt.Println(current())
	}
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
