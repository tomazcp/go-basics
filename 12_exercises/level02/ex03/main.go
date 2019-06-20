package main

import "fmt"

const a = "This is an untyped constant"
const b string = "This is a typed constant"

func main() {
	fmt.Println(a)
	fmt.Println(b)
}
