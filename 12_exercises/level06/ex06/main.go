package main

import "fmt"

func main() {
	myInt := func(x int, y int) int {
		return x * y
	}(5, 10)
	fmt.Println(myInt)
}
