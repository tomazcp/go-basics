package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	newTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	fmt.Println("Truck:", newTruck)

	newSedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println("Sedan:", newSedan)

	fmt.Println(newTruck.color)
	fmt.Println(newSedan.luxury)
}
