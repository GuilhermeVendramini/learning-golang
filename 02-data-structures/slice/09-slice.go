package main

import "fmt"

func main() {

	/*
		3 differents way to declare a slice:

		var students [][]string
		students := [][]string{}
		students := make([][]string, 3)
	*/

	//best way
	students := make([][]string, 3)

	student := make([]string, 3)
	student[0] = "Guilherme"
	student[1] = "Sophia"
	student[2] = "Malu"
	fmt.Println(student)

	student2 := make([]string, 3)
	student2[0] = "Bruno"
	student2[1] = "Carol"
	student2[2] = "Celina"
	fmt.Println(student2)

	/*
		students[0] = student
		students[1] = student2
		students[2] = student

		The 3 item will give error, so use append
		students[3] = student

	*/

	students = append(students, student, student2, student)

	fmt.Println(students)
}
