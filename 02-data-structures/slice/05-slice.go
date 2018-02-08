package main

import "fmt"

func main() {
	/*
		3 is lenght & capacity
		lenght - number of elements referred to by the slice
		capacity - number of elements in the underlying array
	*/
	customerNumber := make([]int, 3)
	customerNumber[0] = 7
	customerNumber[1] = 20
	customerNumber[2] = 15
	/*
		The item below won't work because the capacity is until 3
		customerNumber[3] = 30

		But doing the following example will works fine.

		customerNumber = append(customerNumber, 30)

		See the file 02-slice.go to more information.
	*/

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])
	//fmt.Println(customerNumber[3])

	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Gonjour!"
	greeting[2] = "Bom dia!"

	fmt.Println(greeting[0])
	fmt.Println(greeting[1])
	fmt.Println(greeting[2])
}
