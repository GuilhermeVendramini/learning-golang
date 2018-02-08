package main

import "fmt"

func main() {
	//Anonymous function
	func() {
		fmt.Println("I'm driving!")
	}() //self execute
}
