package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("Range over a slice, i = %v, v = %v \n", i, v)
	}
}
