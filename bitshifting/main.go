package main

import "fmt"

const (
	_ = iota
	//kb = 1024
	//mb = kb * 1024
	//gb = 1024 * mb
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("%d %b\n", x, x) // %b prints bit representation

	y := x << 1                 // bit shifting from 100 -> 1000
	fmt.Printf("%d %b\n", y, y) // %b prints bit representation

	fmt.Println("###############################")

	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
