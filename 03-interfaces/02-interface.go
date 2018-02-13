package main

import (
	"fmt"
	"math"
)

// Interface never implement behavers, just declare.
type Shape interface {
	area() float64
}

// Type Square
type Square struct {
	side float64
}

// Type Circle
type Circle struct {
	radius float64
}

// Behavers implementation for Square type
func (s Square) area() float64 {
	return s.side * s.side
}

// Behavers implementation for Circle type
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	fmt.Println("Square:")
	s := Square{10}
	info(s)
	fmt.Println("Cricle:")
	c := Circle{10}
	info(c)
}
