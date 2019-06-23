package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
}
