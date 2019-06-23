package main

import "fmt"

func main() {
	//x := type{values} composite literal
	x := []int{4, 5, 7, 42, 11}
	fmt.Println(x)
	// i -> index, v -> value
	for i, v := range x {
		fmt.Println(i, v)
	}

	y := []int{4, 5, 7, 42, 11}
	fmt.Println(y[1:])  // from index 1 to the end
	fmt.Println(y[1:3]) // from index 1 to 3 not included

	y = append(y, 10, 25, 45)
	fmt.Println(y)
	y = append(y, x...) // ... spread operator do javascript
	fmt.Println(y)

	// deleting from a slice
	x = append(x, y...) // adding more values to the slice
	fmt.Println(x)
	x = append(x[:2], x[4:]...) // deleting index 2 and 3 values 7 and 42
	fmt.Println(x)

	// make
	w := make([]int, 10, 12)
	fmt.Println(w)
	fmt.Println(len(w)) // length 10
	fmt.Println(cap(w)) // underlying array 12
	w = append(w, 32)
	w = append(w, 33)
	fmt.Println(w)
	fmt.Println(len(w))
	fmt.Println(cap(w))
	w = append(w, 34)
	fmt.Println(w)
	fmt.Println(len(w))
	fmt.Println(cap(w)) // when max capacity is passed, a new slice is created with double the original size

	// multidimensional slice
	myself := []string{"Tomás", "Pinto", "Nata", "Lemon", "Strawberry"}
	raquel := []string{"Raquel", "Nogueira", "Chocolate", "Melon", "Strat"}

	multidimensional := [][]string{myself, raquel}
	fmt.Println(multidimensional)
}
