package main

import "fmt"

func main() {
	arr := [5]int{42, 43, 44, 45, 46}
	fmt.Printf("arr = %v\n", arr)

	for k, v := range arr {
		fmt.Printf("array: key = %v, value = %v, type = %T\n", k, v, v)
	}
}
