package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double zero seven",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   18,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Person.First, p1.Person.Last, p1.Person.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Person.First, p2.Person.Last, p2.Person.Age, p2.LicenseToKill)
}
