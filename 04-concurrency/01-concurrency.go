/*
	In this example nothing will print
	Because everything (main, foo and bar) happens at the same time
	so we don't have time to see what's happening here
*/

package main

import "fmt"

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
