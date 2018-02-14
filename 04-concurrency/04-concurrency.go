/*
	Parallelism:
	In this example we add runtime.GOMAXPROCS
	so we can have parallelism.
	https://golang.org/pkg/runtime/
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

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
