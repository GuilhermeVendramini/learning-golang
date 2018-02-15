/*
	In comment we have wrong examples that will
	cause deadlock or won't work as expected
	and more bellow you can see the solution
*/

package main

import (
	"fmt"
)

/* Wrong way
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c)
}
*/

// Solution
func main() {
	n := 10
	c := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
	}()

	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
}
