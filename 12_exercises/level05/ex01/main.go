package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	icecreamFlavors []string
}

func main() {
	p1 := person{
		firstName:       "Tomás",
		lastName:        "Pinto",
		icecreamFlavors: []string{"Morango", "Nata", "Limão"},
	}
	p2 := person{
		firstName:       "Raquel",
		lastName:        "Nogueira",
		icecreamFlavors: []string{"Chocolate", "Strat", "Baunilha"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	for _, v := range p1.icecreamFlavors {
		fmt.Println(v)
	}
	for _, v := range p2.icecreamFlavors {
		fmt.Println(v)
	}

	// ex02
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	for _, v := range m {
		fmt.Println(v.firstName, v.lastName)
		for _, s := range v.icecreamFlavors {
			fmt.Print(s, ",")
		}
		fmt.Println("")
	}
}
