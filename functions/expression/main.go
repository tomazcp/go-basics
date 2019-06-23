package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	a := func(s string) {
		fmt.Println(s)
	}
	a("I was created to print this")

	myFunc := bar()
	fmt.Printf("%T\n", myFunc) // func() int
	myInt := myFunc()
	fmt.Println(myInt)

	greetFunc := buildGreeting("Tomás")
	greet := greetFunc()
	fmt.Println(greet)

	myCounter := counter(0)
	myCounter()
	myCounter()
	myCounter()
	myCounter()
}

// function that returns a function
func bar() func() int {
	i := 451
	return func() int {
		return i
	}
}

// closure is a function which has access to the variable from another function's scope
func buildGreeting(name string) func() string {
	greeting := fmt.Sprintf("Hello, %v", name)
	return func() string {
		return fmt.Sprintf("%v, Welcome!", greeting)
	}
}

func counter(start int) func() {
	count := start
	return func() {
		fmt.Println(count)
		count++
	}
}
