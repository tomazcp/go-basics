package main

import (
	"fmt"
	"sort"
)

// Person some text
type Person struct {
	First string
	Age   int
}

// ByAge some text
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i int, j int) bool {
	return a[i].Age < a[j].Age
}

// ByFirst some text
type ByFirst []Person

func (a ByFirst) Len() int {
	return len(a)
}

func (a ByFirst) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByFirst) Less(i, j int) bool {
	return a[i].First < a[j].First
}

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	isSorted := sort.IntsAreSorted(s)
	fmt.Println(isSorted)
	sort.Ints(s)
	fmt.Println(s)
	isSorted = sort.IntsAreSorted(s)
	fmt.Println(isSorted)

	p1 := Person{
		First: "Tomás",
		Age:   32,
	}

	p2 := Person{
		First: "Raquel",
		Age:   23,
	}

	p3 := Person{
		First: "MJ",
		Age:   62,
	}

	p4 := Person{
		First: "Jaime",
		Age:   27,
	}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByFirst(people))
	fmt.Println(people)
}
