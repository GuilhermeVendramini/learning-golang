package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"Bom dia!",
		"Bongiorno",
		"Ohayo!",
		"Gutten morgen",
		"Buenos dias!",
	}

	fmt.Print("[1:2]")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2]")
	fmt.Println(greeting[:2])
	fmt.Print("[5:]")
	fmt.Println(greeting[5:])
	fmt.Print("[:]")
	fmt.Println(greeting[:])
	fmt.Print("greeting")
	fmt.Println(greeting)
}
