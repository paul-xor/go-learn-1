package main

import "fmt"

func main() {
	var number uint = 2
	const img = "lalal"
	fmt.Printf("Hello Gopher! %v\n", img)
	fmt.Println(number)

	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)
}
