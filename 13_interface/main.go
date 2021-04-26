package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.h * r.w
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c1 := Circle{10}
	r1 := Rectangle{10, 10}

	fmt.Printf("Circle Area : %f\n", getArea(c1))
	fmt.Printf("Rectangle Area : %f\n", getArea(r1))
}
