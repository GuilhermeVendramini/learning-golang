package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}

	//To join 2 slices use ... at the end
	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], myOtherSlice[1:]...)
	fmt.Println(mySlice)
}
