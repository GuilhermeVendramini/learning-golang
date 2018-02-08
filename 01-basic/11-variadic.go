package main

import "fmt"

var g_groceries []string

func main() {
	add_grocery("Banana", "Apple", "Bread", "Milk", "Potato", "Salt", "Mango")
	list_groceries()
	fmt.Println("Capacity", cap(g_groceries))
}

func add_grocery(a ...string) {
	for _, d := range a {
		g_groceries = append(g_groceries, d)
	}
}

func list_groceries() {
	fmt.Println("Grocery list:")

	for _, d := range g_groceries {
		fmt.Println(d)
	}
}
