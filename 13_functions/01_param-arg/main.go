package main

import "fmt"

func main() {
	greet("Tomaz") // argument "Tomaz"
	greet("Jaime") // argument "Jaime"
	printFullName("Tomaz", "Pinto")
}

// @param {String} name
func greet(name string) {
	fmt.Println("Hello", name)
}

// when two or more params are of the same type
// we only need to declare the type on the last param
func printFullName(fname, lname string) {
	fmt.Println(fname, lname)
}
