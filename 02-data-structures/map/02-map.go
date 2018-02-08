package main

import "fmt"

func main() {

	/*
		Three ways to create a map:
		var myGreeting = map[string]string{}
		myGreeting := map[string]string{}
		var myGreeting = make(map[string]string)
	*/
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}
