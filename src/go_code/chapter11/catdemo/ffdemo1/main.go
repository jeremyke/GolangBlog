package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var circle1 Circle
	circle1.radius = 4.0
	area := circle1.area()
	fmt.Println(area)
}
