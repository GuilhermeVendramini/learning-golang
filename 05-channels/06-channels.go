/*
	We can also use one channel in more than one
	function at the same time.
	So we can have more than one function executing the same channel
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	//Executing the channel c
	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	//Executing the channel c again
	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	<-done
	<-done
}
