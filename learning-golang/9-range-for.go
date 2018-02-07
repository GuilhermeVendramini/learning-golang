package main

import "fmt"

var g_groceries []string

func main() {
	add_grocery("Banana")
	add_grocery("Apple")
	add_grocery("Bread")
	add_grocery("Milk")
	add_grocery("Potato")
	add_grocery("Salt")
	add_grocery("Mango")
	list_groceries()
}

func add_grocery(a string) {
	fmt.Println("Capacity", cap(g_groceries))
	g_groceries = append(g_groceries, a)
}

func list_groceries() {
	fmt.Println("Grocery list:")
	/*for i := 0; i < len(g_groceries); i++ {
		fmt.Println(g_groceries[i])
	}*/

	/*for i, d := range g_groceries {
		fmt.Println(i, d)
	}*/

	for _, d := range g_groceries {
		fmt.Println(d)
	}
}
