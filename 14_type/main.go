package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	var b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b) // em go cast é converção
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
