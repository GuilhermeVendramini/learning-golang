package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	/*
		Without key
		for _, val := range myGreeting {
			fmt.Println("Val: ", val)
		}
	*/

	//With key
	for key, val := range myGreeting {
		fmt.Println("Key:", key, "Val: ", val)
	}
}
