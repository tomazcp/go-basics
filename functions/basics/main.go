package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	foo("Go")
	p1 := person{
		firstname: "Tomás",
		lastname:  "Pinto",
	}
	changeName(p1)
	// firstname and lastname unchanged for p1, pass by value
	fmt.Println(p1.firstname, p1.lastname)

	str := bar("bar")
	fmt.Println(str)

	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	sliceOfInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := "some string"
	// the variadic param has always to be the last parameter
	sum := sum(s)
	fmt.Println(sum)

	result := multiplyAll(sliceOfInt...)
	fmt.Println(result)
}

// everything in go is pass by value
func foo(s string) {
	fmt.Println("Hello,", s)
}

func changeName(p person) {
	p.firstname = "Firstname change"
	p.lastname = "Lastname change"
}

// returns string
func bar(s string) string {
	return fmt.Sprintf("Hello from woo, %v", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}

// x is stored as a slice of the type
func sum(s string, x ...int) int {
	fmt.Println("##############")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("##############")

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func multiplyAll(x ...int) int {
	result := 1
	for _, v := range x {
		result *= v
	}
	return result
}
