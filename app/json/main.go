package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := person{
		First: "Tomás",
		Last:  "Pinto",
		Age:   32,
	}

	p2 := person{
		First: "Raquel",
		Last:  "Nogueira",
		Age:   23,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	// marshal
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// unmarshal
	var newPeople []person
	data := `[{"First":"Tomás","Last":"Pinto","Age":32},{"First":"Raquel","Last":"Nogueira","Age":23}]`
	err2 := json.Unmarshal([]byte(data), &newPeople)
	if err2 != nil {
		fmt.Println(err)
	}
	for i, v := range newPeople {
		fmt.Println("--- PRINTING PERSON NUMBER", i+1)
		fmt.Println(v)
	}
}

type person struct {
	First string
	Last  string
	Age   int
}
