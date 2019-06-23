package main

import "fmt"

func main() {
	// ex01
	a := 42
	fmt.Println(&a)

	// ex02
	p := person{
		first: "Tomás",
		last:  "Pinto",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "Raquel"  // (*p).first
	p.last = "Nogueira" // (*p).last
}
