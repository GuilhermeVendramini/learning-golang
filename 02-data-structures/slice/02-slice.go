package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("------------")

	/*The slice capacity is 5 but in this case the loop is going until 80.
	  So a new slice will be created duplicating the capacity until satisfy 80.
	  And the old slice will be replaced by the new one.
	*/
	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}
}
