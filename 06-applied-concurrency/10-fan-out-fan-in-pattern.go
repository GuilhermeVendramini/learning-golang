package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 5, 6)

	//FAN OUT
	/*
		Distribute the sq word across two goroutines that both read from "in"
		So here we have 2 "sq" to process 1 "in"
	*/
	c1 := sq(in)
	c2 := sq(in)

	//FAN IN
	//Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

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

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	/*
		Start an output goroutine for each input channel in cs.
		Output copies values from c to out until c is closed, then calls wg.Done.
	*/
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	/*
		Start a goroutine to close out once all the output goroutines are done.
		This must start after the wg.Add call.
	*/
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
FAN OUT
Multiple functions reading from the same
channel until that channel is closed.

FAN IN
A function can read from multiple inputs
and proceed until all are closed by multiplexing
the input channels onto a single channel that's
closed when all the inputs are closed.

Source:
https://blog.golang.org/pipelines
*/
