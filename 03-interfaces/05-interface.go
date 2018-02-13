/*
	In this example we don't need implement "Len", "Swap" and "Less".
	Because we're using "StringSlice" and in "StringSlice" was already implemented.
	https://golang.org/pkg/sort/#StringSlice
*/

package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	studyGroup := people{"Gui", "Carol", "Ka", "Sophia"}

	fmt.Println(studyGroup)
	sort.StringSlice(studyGroup).Sort()
	/*
		Another way to use:
		sort.Sort(sort.StringSlice(studyGroup))
	*/
	fmt.Println(studyGroup)
}
