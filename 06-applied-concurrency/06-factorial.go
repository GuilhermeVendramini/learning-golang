/*
	Here we have 2 examples:
	With and without channels
*/

package main

import (
	"fmt"
)

/* Without go and channel
func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
*/

// With go and channel
func main() {
	for f := range factorial(4) {
		fmt.Println("Total:", f)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
