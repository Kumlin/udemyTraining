package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return (c.radius * c.radius) * math.Pi
}

func (s Square) area() float32 {
	return s.side * s.side
}

type shape interface {
	area() float32
}

func shapeArea(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	myCircle := Circle{10}
	mySquare := Square{5}

	shapeArea(myCircle)
	shapeArea(mySquare)
}
