package main

import (
	"fmt"
	"log"
)

type MathError struct {
	lat, long string
	err       error
}

func (n *MathError) Error() string {
	return fmt.Sprintf("Math message error: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		me := fmt.Errorf("Specifc message error: %v", f)
		return 0, &MathError{"23.23 N", "92.23 W", me}
	}

	return f, nil
}
