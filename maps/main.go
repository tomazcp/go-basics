package main

import "fmt"

func main() {
	// composite literal map of key type string and values int
	m := map[string]int{
		"tomaz":  32,
		"raquel": 23,
	}
	fmt.Println(m)
	fmt.Println(m["tomaz"])
	// if the key doesn't exist
	fmt.Println(m["mom"]) // 0

	v, ok := m["mom"]
	fmt.Println(v)  // 0
	fmt.Println(ok) // false
	// in idiomatic go -> comma ok idiom
	if v, ok := m["mom"]; ok {
		fmt.Printf("it was ok %v\n", v)
	} else {
		fmt.Printf("It wasn't ok %v\n", v)
	}

	m["jaime"] = 27

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "jaime")
	fmt.Println(m)
}
