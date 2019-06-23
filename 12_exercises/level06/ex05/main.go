package main

import (
	"fmt"
	"math"
)

func main() {
	s := rectangle{
		length: 10,
		width:  10,
	}

	c := circle{
		radius: 10,
	}

	info(s)
	info(c)
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width  float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println("The area of this circle is", s.area())
	case rectangle:
		fmt.Println("The area of this rectangle is", s.area())
	}
}
