package main

import "fmt"

func main() {
	m := map[string][]string{
		"tomaz":  {"coding", "dogs", "food"},
		"raquel": {"bac", "gatos", "viajar"},
	}
	for k, v := range m {
		fmt.Println("This is for record", k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
	// ex09
	m["jaime"] = []string{"motas", "anime", "gatos"}

	for k, v := range m {
		fmt.Println("This is for record", k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
	// ex10
	delete(m, "jaime")
	fmt.Println(m)
}
