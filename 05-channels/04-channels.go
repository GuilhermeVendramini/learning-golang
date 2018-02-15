package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			//set c value and waiting to be used
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		//Close the channels
		<-done
		<-done
		//Close c
		close(c)
	}()

	/*
		Here c is used, so after that c can
		be processed by the "for" above again
	*/
	for n := range c {
		fmt.Println(n)
	}
}
