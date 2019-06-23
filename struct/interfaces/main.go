package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
	last  string
}

func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first, sa.last, "- secret agent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- person speak")
}

type secretAgent struct {
	person
	ltk bool
}

func bar(h human) {
	// switch on type
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
	fmt.Println(h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Tomás",
			last:  "Pinto",
		},
		ltk: false,
	}

	sa2 := secretAgent{
		person: person{
			first: "Raquel",
			last:  "Nogueira",
		},
		ltk: true,
	}

	sa1.speak()
	sa2.speak()

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	bar(sa1)
	bar(sa2)
	bar(p1)

}
