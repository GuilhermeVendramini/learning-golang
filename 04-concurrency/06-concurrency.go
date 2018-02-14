/*
	Mutex example.
	In this example we're locking the concurrency
	making the "counter" and after that unlocking the concurrency
	so we can have a precise counter
	http://pt.wikipedia.org/wiki/Mutex
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
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		mutex.Lock() //Lock so we can count
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
