package main

import "fmt"

func main() {
	c := gen(2, 3)
	out := sq(c)

	for o := range out {
		fmt.Println(o)
	}
	/*	We can also print 2 times instead of "for"
		fmt.Println(<-out)
		fmt.Println(<-out)
	*/
}

/*
	First - return chan = 2
	So "sq" is called to process chan = 2 (2 * 2 = 4)
	and "main" print chan = 2 (out 4).

	Second - return chan = 3
	So "sq" is called to process chan = 3 (3 * 3 = 9)
	and "main" print chan = 3 (out 9).
*/
func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
