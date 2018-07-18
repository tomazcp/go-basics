package main

import "fmt"

func main() {
	for i := 5000; i <= 5500; i++ {
		// print numeric value i
		// convert the numeric value into a string and print
		// convert the string into bytes
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
