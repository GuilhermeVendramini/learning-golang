package main

import "fmt"

//total capacity of our list
const g_cap int = 5

var g_groceries [g_cap]string

//total number of grocery items in our current array
var g_len int = 0

func main() {
	add_grocery("Banana")
	add_grocery("Apple")
	add_grocery("Bread")
	add_grocery("Milk")
	add_grocery("Potato")
	add_grocery("Salt")
	list_groceries()
}

func add_grocery(a string) {
	if g_len < g_cap {
		g_groceries[g_len] = a
		g_len++
	} else {
		fmt.Println("Too many items.")
	}
}

func list_groceries() {
	fmt.Println("Grocery list:")
	for i := 0; i < g_len; i++ {
		fmt.Println(g_groceries[i])
	}
}
