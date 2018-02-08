package main

import (
	"fmt"
)

func main() {
	b := true

	/*
		Initialized food at the if statement
		So food can be used just  inside the if statement
	*/
	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	/*
		That won't work because food is out of if statement
		fmt.Println(food)
	*/

}
