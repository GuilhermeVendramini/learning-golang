package main

import "fmt"

func main() {
	for n := range sq(gen(2, 3)) {
		fmt.Println(n)
	}
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
