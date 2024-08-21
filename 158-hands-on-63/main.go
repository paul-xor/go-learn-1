package main

import "fmt"

func main() {
	fmt.Println(Add(3, 4))
}

func add(a int, b int) int {
	return a + b
}
