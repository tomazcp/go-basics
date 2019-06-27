package main

import "fmt"

func main() {
	// buffered channel não vai bloquear
	// porque o channel aceita dois inteiros
	ch := make(chan int, 2)
	ch <- 43
	ch <- 44
	fmt.Println(<-ch)
	fmt.Println("hello")
	fmt.Println(<-ch)
}
