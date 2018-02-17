package main

import (
	"errors"
	"log"
)

//Start a error variable  with "Err" is a good practice
var ErrMessage = errors.New("Error message here!")

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMessage
	}

	return f, nil
}
