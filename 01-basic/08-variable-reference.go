package main

import "fmt"

/*
	& = referemce
	* = de reference
*/

//&c Send by reference
func main() {
	var c int
	fmt.Println(&c)
	max(100, 15, &c)
	fmt.Println(c)
}

//*k assume the referemce of &c
func max(i int, j int, k *int) {
	fmt.Println(k)
	if i > j {
		*k = i
	} else {
		*k = j
	}
}
