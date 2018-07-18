package main

import (
	"fmt"
)

func zero(z *int) {
	fmt.Println(z) // address 0xc420016088
	*z = 0
}

func main() {
	x := 5
	fmt.Println(&x) // address 0xc420016088
	zero(&x)
	fmt.Println(x) // x is 0
}
