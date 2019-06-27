package main

import "fmt"

func main() {
	// vai bloquear porque só alocamos espaço
	// no buffer para um valor
	ch := make(chan int)
	ch <- 43
	ch <- 44 // block
	fmt.Println(<-ch)
	fmt.Println("hello")
}
