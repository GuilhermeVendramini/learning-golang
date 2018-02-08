package main

import "fmt"

func main() {
	for day := 1; day <= 5; day++ {
		fmt.Println("On the", day, "of Christmas my friend send to me:")
		switch day {
		case 1:
			fmt.Println("Gift 1")
		case 2:
			fmt.Println("Gift 2")
		case 3:
			fmt.Println("Gift 3")
		case 4:
			fmt.Println("Gift 4")
		default:
			fmt.Println("Another Gift")
		}
	}
}
