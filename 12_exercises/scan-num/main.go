package main

import "fmt"

func main() {
	var first int
	var second int
	remainder := 0

	fmt.Println("Enter a value:")
	fmt.Scan(&first)
	fmt.Println("Enter a second value:")
	fmt.Scan(&second)

	remainder = calculateRemainder(first, second)

	fmt.Printf("Remainder is %v\n", remainder)
}

func calculateRemainder(first int, second int) int {
	if first > second {
		return (first % second)
	}
	return second % first
}
