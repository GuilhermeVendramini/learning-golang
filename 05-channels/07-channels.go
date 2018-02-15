/*
	Channel as Arguments Returns
*/

package main

import (
	"fmt"
)

func main() {
	c := incrementor()
	for n := range puller(c) {
		fmt.Println("Out: ", n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

/*
	sum += n
	0 + 0 = 0
	0 + 1 = 1
	1 + 2 = 3
	3 + 3 = 6
	Out = 6
*/
func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			fmt.Println(n)
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
