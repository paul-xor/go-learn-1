package main

import "fmt"

func main() {
	f := foo()
	i, s := bar()
	fmt.Printf("foo returns: %v\t bar returns: %v, %v\n", f, i, s)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	i := 42
	s := "James Bond"
	return i, s
}
