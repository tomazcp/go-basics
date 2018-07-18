package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64 // declare var
	fmt.Print("Enter meters swam \n")
	fmt.Scan(&meters) // here we need to use memory address of meters
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards")
}
