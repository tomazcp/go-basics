package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 25)
	c := (74 >= 100)
	d := (35 != 27)
	e := (32 < 32)
	f := (32 > 32)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
