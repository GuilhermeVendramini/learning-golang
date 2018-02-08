package main

import "fmt"

func main() {

	/*
		Two ways to create a map:
		var myGreeting = map[string]string{}
		var myGreeting = make(map[string]string)
	*/
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}
