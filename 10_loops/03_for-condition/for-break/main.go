package main

import "fmt"

func main() {
	i := 0
	// like while but without condition
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}
