package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) fullName() string {
	return p.First + p.Last
}

func (p Person) getAge() int {
	return p.Age
}

func main() {
	p1 := Person{"James", "Bond", 30}
	p2 := Person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p1.getAge())
	fmt.Println(p2.fullName())
	fmt.Println(p2.getAge())
}
