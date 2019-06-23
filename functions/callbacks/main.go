package main

import "fmt"

func main() {
	myInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sum(myInts...)
	fmt.Println(s)

	evensSum := even(sum, myInts...)
	fmt.Println(evensSum)

	oddsSum := odd(sum, myInts...)
	fmt.Println(oddsSum)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, values ...int) int {
	var yi []int
	for _, v := range values {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(x ...int) int, values ...int) int {
	var allValues []int
	for _, v := range values {
		if v%2 != 0 {
			allValues = append(allValues, v)
		}
	}
	return f(allValues...)
}
