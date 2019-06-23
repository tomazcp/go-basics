package main

import "fmt"

func main() {
	p1 := person{
		first: "Tomás",
		last:  "Pinto",
		age:   32,
	}
	p1.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I'm %v years old.\n", p.first, p.last, p.age)
}
