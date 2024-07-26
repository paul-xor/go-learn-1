package main

import "fmt"

func main() {
	x := []int{3, 4, 5, 6, 7}
	sumf := foo(x...)
	sumb := bar(x)
	fmt.Println(sumf, sumb)
}

func foo(ii ...int) int {
	fmt.Printf("%T\n", ii)
	sum := 0
	for i, v := range ii {
		fmt.Printf("index: %v, value: %v\n", i, v)
		sum += v
	}
	return sum
}

func bar(ii []int) int {
	fmt.Printf("%T\n", ii)
	sum := 0
	for i, v := range ii {
		fmt.Printf("index: %v, value: %v\n", i, v)
		sum += v
	}
	return sum
}
