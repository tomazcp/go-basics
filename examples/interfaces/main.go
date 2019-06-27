package main

import "fmt"

type person struct {
	first string
	human
}

type human interface {
	speak()
	sleep()
}

func (p *person) speak() {
	fmt.Println(p.first, "says hello")
}

func main() {
	p := new(person)
	p.first = "Tomás"
	fmt.Printf("%T\n", p)
	triggerActions(p)
}

func triggerActions(h human) {
	h.speak()
}
