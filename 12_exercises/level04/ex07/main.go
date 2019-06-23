package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "Hellooooooo, James"}

	people := [][]string{jb, mm}
	fmt.Println(people)
	for _, v := range people {
		for j, k := range v {
			fmt.Println(j, k)
		}
	}

}
