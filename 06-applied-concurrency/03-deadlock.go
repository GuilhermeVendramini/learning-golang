/*
	In comment we have wrong examples that will
	cause deadlock and more bellow you can see the solution
*/

package main

import (
	"fmt"
)

/* Wrong way
func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}
*/

// Solution
func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()

	fmt.Println(<-c)
}
