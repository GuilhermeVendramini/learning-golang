/*
	Race conditions example.
	In this example we're using sleep time random
	so we can have the "counter" with the same value sometimes
	because we aren't locking the concurrency
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i <= 5; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
