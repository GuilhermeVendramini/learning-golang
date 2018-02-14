package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

/*
	Override the default string print.
	https://golang.org/doc/effective_go.html#printing
*/
func (p person) String() string {
	return fmt.Sprintf("People: %s: %d", p.Name, p.Age)
}

type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Gui", 28},
		{"Carol", 25},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
