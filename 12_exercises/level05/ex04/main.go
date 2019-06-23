package main

import "fmt"

func main() {
	anon := struct {
		name string
		age  int
	}{
		name: "John Doe",
		age:  0,
	}

	fmt.Println(anon)
}
