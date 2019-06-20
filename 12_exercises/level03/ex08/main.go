package main

import "fmt"

func main() {
	name := "MJ"
	switch {
	case name == "Raquel":
		fmt.Println("namorada")
	case name == "MJ":
		fmt.Println("mãe")
	case name == "jaime":
		fmt.Println("irmão")
	}
}
