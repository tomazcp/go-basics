package main

import "fmt"

// bill kennedy: maioria das vezes usar unbuffered channels
// ch := make(chan type)

func main() {
	// does not work because channels block
	// exemplo corrida de estafetas, só quando é possível
	// transmitir o testemunho (send & receive) ao mesmo tempo
	// é que o channel desbloqueia (successful passage of value)
	// c := make(chan int)
	// c <- 42
	// fmt.Println(<-c)
	c := make(chan int)
	// código bloqueia nesta goroutine
	// mas sendo que não está a correr na main
	// goroutine não bloqueia, logo é possível aceder
	// ao valor fora da goroutine
	go func() {
		c <- 42
	}()
	fmt.Println(<-c) // esta bloqueia até ser passado o valor
}
