package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", slice)

	// ex03
	fmt.Println(append(slice[:5]))
	fmt.Println(append(slice[5:]))
	fmt.Println(append(slice[2:7]))
	fmt.Println(append(slice[1:6]))
}
