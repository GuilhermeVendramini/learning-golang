/*
	We add WaitGroup, so we can see what's happens
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//add 2 Wait Group
	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	//Here we done the first group
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	//Here we done the second group
	wg.Done()
}
