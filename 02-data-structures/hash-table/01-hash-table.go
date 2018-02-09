package main

import "fmt"

func main() {
	/*
		When I use '' this means letter
		When I use "" this means string

		I can also do this:
		letter := rune("A"[0])
	*/
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	/*
		So we can use ASCII like hash.
		See the example 02-hash-table.go
		Look https://pt.wikipedia.org/wiki/ASCII
	*/
}
