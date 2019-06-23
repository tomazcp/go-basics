package main

import "fmt"

// type person with underlying type struct
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	// anonymous struct
	anonStruct := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   43,
	}
	fmt.Println(anonStruct)

}
