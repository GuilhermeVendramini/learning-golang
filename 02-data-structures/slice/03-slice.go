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

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}
