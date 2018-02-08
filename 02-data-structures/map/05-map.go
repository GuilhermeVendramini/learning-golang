package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)
	//delete(myGreeting, 2)
	//fmt.Println(myGreeting)

	/*
		val //Get the condition value. That also can be _
		exists := myGreeting[2] //Condition
		exists //Get condition value True or False
	*/
	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	}

	//Add a new item
	myGreeting[4] = "Bom dia!"

	fmt.Println(myGreeting)
}
