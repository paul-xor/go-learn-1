package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 4))
}

func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v", a, x)
}

func square(n int) int {
	return n * n
}
