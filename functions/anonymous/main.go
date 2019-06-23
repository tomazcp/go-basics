package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println(x, "was passed into the function")
	}(42)
}
