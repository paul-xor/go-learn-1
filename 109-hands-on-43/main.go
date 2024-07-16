package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for k, v := range xi {
		fmt.Printf("k = %v,\t v = %v \n", k, v)
	}

	fmt.Printf("%T \t %#v\n", xi, xi)
}
