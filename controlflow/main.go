package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough // if it executes the code inside the case, all subsquent case statements will be executed, even if they're false
	case (4 == 4):
		fmt.Println("also true does it print?")
	}
}
