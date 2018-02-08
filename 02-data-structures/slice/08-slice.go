package main

import "fmt"

func main() {
	student := make([]string, 3)
	students := make([]string, 35)
	student[0] = "Guilherme"
	student[1] = "Sophia"
	student[2] = "Malu"
	students = append(student, "Carol")

	fmt.Println(students)
	fmt.Println("Len:", len(students), "Capacity:", cap(students))
	fmt.Printf("%v %T \n", students[2], students[2])
	fmt.Println("---------------")
	students = append(students, "Celina", "Bruno", "Maike")
	fmt.Println(students)
	fmt.Println("Len:", len(students), "Capacity:", cap(students))
}
