/*
	We add time.Sleep
	So we can see foo running faster than Bar
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 8; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(1 * time.Second))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 8; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(2 * time.Second))
	}
	wg.Done()
}
