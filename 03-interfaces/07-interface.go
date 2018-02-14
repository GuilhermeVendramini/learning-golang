package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{7, 4, 2, 9, 10, 1, 23, 3}

	fmt.Println(s)
	/*
		Other ways to use:
		sort.Ints(s)
	*/
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)

}
