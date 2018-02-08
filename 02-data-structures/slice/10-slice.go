package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 10
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])
}
