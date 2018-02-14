package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func (v vehicle) Specs() {
	fmt.Printf("Seats %v, Speed %v, color %v\n", v.Seats, v.MaxSpeed, v.Color)
}

func main() {
	//cars
	prius := car{}
	tacoma := car{}

	//plane
	boeing747 := plane{}
	boeing757 := plane{}

	//boat
	sanger := boat{}
	nautique := boat{}

	rides := []vehicles{
		prius,
		tacoma,
		boeing747,
		boeing757,
		sanger,
		nautique,
	}

	for k, v := range rides {
		fmt.Println(k, " - ", v)
	}
}
