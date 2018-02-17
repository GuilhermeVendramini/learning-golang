package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)

		//Differents ways to show a error
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}
