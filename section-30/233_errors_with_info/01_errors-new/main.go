package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	sq, err := sqrt(-2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sq)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return f * f, nil
}
