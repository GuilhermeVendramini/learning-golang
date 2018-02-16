package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	f := fanIn(boring("Ma"), boring("Ka"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-f)
	}
	fmt.Println("You're both boring.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(c1, c2 <-chan string) <-chan string {
	f := make(chan string)
	go func() {
		for {
			f <- <-c1
		}
	}()

	go func() {
		for {
			f <- <-c2
		}
	}()
	return f
}
