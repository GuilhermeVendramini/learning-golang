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
		Other ways to use:
		sort.Sort(sort.StringSlice(studyGroup))
		sort.Strings(studyGroup)
	*/
	fmt.Println(studyGroup)
}
