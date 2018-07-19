package main

import "fmt"

func main() {
	// at this point the args are of type "kind"
	// the type is undetermined
	n := average(43, 57, 9, 14, 21, 1037)
	fmt.Println(n)

	// example 2
	data := []float64{43, 57, 9, 14, 21, 1037} // slice of float64
	r := average2(data...)                     // passing slice as individual values to the func
	fmt.Println(r)
}

// variadic function
// with ... we can pass an arbitrary number of params
func average(sf ...float64) float64 {
	fmt.Println(sf)        // [arg1, arg2, etc..] => slice of args
	fmt.Printf("%T\n", sf) // []float64 => slice float64
	// total := 0.0
	var total float64 // declaring this way sets total to the zero value automatically
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func average2(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
