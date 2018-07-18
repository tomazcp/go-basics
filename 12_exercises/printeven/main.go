package main

import "fmt"

func main() {
	max := 100
	printEvens(max)
}

func printEvens(max int) {
	for i := 1; i <= max; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
