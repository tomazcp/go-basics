package main

import "fmt"

func main() {
	a := 11
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	a = a << 1
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

}
