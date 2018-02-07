package main

import "fmt"

func main() {
	age := 20

	if age < 13 {
		fmt.Println("You're young!")
	} else if age > 12 && age < 20 {
		fmt.Println("You're a teenager")
	} else {
		fmt.Println("You're an adult")
	}
}
