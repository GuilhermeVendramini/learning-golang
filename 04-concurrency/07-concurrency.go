/*
	Atomic
	Using Atomic we won't have problems with the counter
	but we can have different sequence
	https://golang.org/pkg/sync/atomic/
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

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
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
