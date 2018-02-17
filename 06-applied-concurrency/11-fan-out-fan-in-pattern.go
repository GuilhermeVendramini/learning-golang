/*
	Atomic
	Using Atomic we won't have problems with the counter
	but we can have different sequence
	https://golang.org/pkg/sync/atomic/
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {

	in := gen("Foo:", "Side", "Bar")

	//FAN OUT
	ic1 := incrementor(in)
	ic2 := incrementor(in)

	//FAN IN
	for n := range merge(ic1, ic2) {
		fmt.Println(n)
	}

	fmt.Println("Total:", counter)
}

func gen(items ...string) chan string {
	out := make(chan string)
	go func() {
		for _, i := range items {
			out <- i
		}
		close(out)
	}()
	return out
}

func incrementor(item chan string) chan string {
	out := make(chan string)
	go func() {
		for i := range item {
			for c := 0; c <= 5; c++ {
				out <- string(i) + " Counter: " + strconv.Itoa(c)
				counter++
			}
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	/*
		Start an output goroutine for each input channel in cs.
		Output copies values from c to out until c is closed, then calls wg.Done.
	*/
	output := func(c <-chan string) {
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
