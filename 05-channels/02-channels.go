package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		/*
			After close the channel we can't put more values
			but we can use the channel
		*/
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
