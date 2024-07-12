package main

import "fmt"

func main() {
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v \n", x, y)

	// Loops / Interative
	// for statements
	// https://go.dev/doc/effective_go#for

	for i := 0; i < 5; i++ {
		fmt.Printf("countinuing to five: %v first for loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second loop\n", y)
		y++
	}

	// break
	//takes out of loop
	for {
		fmt.Printf("y is %v \t\t\t third loop\n", y)
		if y > 20 {
			break
		}
		y++
	}
	// continue
	// takes to the next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("continue even numbers: ", i)
	}
}
