package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("---")
		for j := 10; j > 5; j-- {
			fmt.Printf("outer loop %v \t inner loop %v \n", i, j)
		}
	}
}
