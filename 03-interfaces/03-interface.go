/*
	In this example we need implement "Len", "Swap" and "Less".
	https://golang.org/pkg/sort/#Interface
*/

package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	studyGroup := people{"Gui", "Carol", "Ka", "Sophia"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println("-----Methods-----")
	fmt.Println(studyGroup.Len())
}
