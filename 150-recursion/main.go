package main

import "fmt"

func main() {
	// 4!
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(5))
	fmt.Println(factLoop(5))
}

// func (r reciever) identifier(p paramenter(s)) (r return(s)) { code }(argument(s))

func factorial(n int) int {
	fmt.Println("this is n: ", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
