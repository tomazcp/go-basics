package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x, y, z)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 5, "i am a string"
}
