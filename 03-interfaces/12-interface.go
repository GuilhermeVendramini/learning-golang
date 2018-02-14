package main

import "fmt"

type Square struct {
	side float64
}

type Shape interface {
	area() float64
}

func (z *Square) area() float64 {
	return z.side * z.side
}

func info(z Shape) {
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(&s)
}
