package main

import "fmt"

func hello() {
	fmt.Println("Hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello() //hello will print first than world
}
