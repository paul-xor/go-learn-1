package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("4 + 6 = ", mySum(4, 6))
	fmt.Println("5 + 9 = ", mySum(5, 9))
	fmt.Println("2 + 10 = ", mySum(2, 10))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
