package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				//set c value and waiting to be used
				c <- i
			}
			//Open the channels
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
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
