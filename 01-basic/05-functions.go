package main

import "fmt"

func main() {
	fmt.Println(max(100, 15))
}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
