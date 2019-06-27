package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// ex01
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Goroutine one launched")
		wg.Done()
	}()

	go func() {
		fmt.Println("Goroutine two launched")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Program exiting...")

	// ex02
	p := person{
		first: "John Doe",
	}
	//p2 := new(person) // returns pointer to person
	//saySomething(p) // can't pass value
	saySomething(&p) // can pass pointer

	// ex03
	// counter := 0
	// var wg2 sync.WaitGroup
	// const delta = 50
	// wg2.Add(delta)
	// for i := 0; i < delta; i++ {
	// 	go func() {
	// 		v := counter
	// 		runtime.Gosched()
	// 		v++
	// 		counter = v
	// 		wg2.Done()
	// 	}()
	// }
	// wg2.Wait()
	// fmt.Println(counter)
	// fmt.Println("program about to exit")

	// ex04
	counter2 := 0
	var mux sync.Mutex
	var wg3 sync.WaitGroup
	delta2 := 50
	wg3.Add(delta2)
	for i := 0; i < delta2; i++ {
		go func() {
			mux.Lock()
			v := counter2
			v++
			counter2 = v
			mux.Unlock()
			wg3.Done()
		}()
	}
	wg3.Wait()
	fmt.Println(counter2)
	fmt.Println("program about to exit")

	// ex05
	var counter3 int64
	const gs = 100
	var wg4 sync.WaitGroup
	wg4.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter3, 1)
			runtime.Gosched()
			fmt.Println("Counter3:", atomic.LoadInt64(&counter3))
			wg4.Done()
		}()
	}
	wg4.Wait()
	fmt.Println("Program about to exit.. ex05")
}

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.first, "is speaking")
}

func saySomething(h human) {
	h.speak()
}
