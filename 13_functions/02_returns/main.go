package main

import "fmt"

func main() {
	// space added because Sprint
	// concatenates strings
	s := greet("Tomaz ", "Pinto")
	fmt.Println(s)

	fmt.Println(greet3("Tomaz ", "Pinto"))
}

// normal return
func greet(fname, lname string) string {
	// Sprint prints to a string (not stout)
	// concatenates strings
	return fmt.Sprint(fname, lname)
}

// named return
func greet2(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname) // assign named return to a value
	return                       // then simply return
}

// return multiple
func greet3(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
