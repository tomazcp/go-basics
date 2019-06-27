package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	const gs = 100

	wg.Add(gs)

	var mux sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mux.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}
