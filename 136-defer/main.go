package main

import "fmt"

func main() {
	defer foo()
	bar()
}

// func (r reciever) identifier(p paremeter(s)) (r return(s)) {code}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
