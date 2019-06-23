package main

import "fmt"

func main() {
	counter := foo()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	f := func(max int) {
		for i := 0; i < max; i++ {
			fmt.Println(i)
		}
	}
	f(100)

	values := []int{16, 32, 24, 65, 2, 7, 11}
	result := evens(multiply, values)
	fmt.Println(result)
}

// closure
func foo() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func evens(cb func(x ...int) int, values []int) int {
	var evenValues []int
	for _, v := range values {
		if v%2 == 0 {
			evenValues = append(evenValues, v)
		}
	}
	return cb(evenValues...)
}

func multiply(values ...int) int {
	total := 1
	for _, v := range values {
		total *= v
	}
	return total
}
