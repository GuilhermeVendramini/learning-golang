package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	//func auto execute
	go func() {
		for i := 0; i < 10; i++ {
			/*
				I set a value to channel c
				The for stop until some place take the channel again
			*/
			c <- i
		}
	}()

	//func auto execute
	go func() {
		for {
			/*
				Here the channel is taken
				So after the print the channel keep going
			*/
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
