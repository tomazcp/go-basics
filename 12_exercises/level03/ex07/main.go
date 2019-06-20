package main

import "fmt"

func main() {
	name := "Raquel"

	if name == "Raquel" {
		fmt.Println("My Baby")
	} else if name == "Jaime" {
		fmt.Println("Irmão")
	} else if name == "Maria João" {
		fmt.Println("Mom")
	} else {
		fmt.Println("Dunno who the fuck you want")
	}
}
