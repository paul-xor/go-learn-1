package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Println("this is an anonymous func showing the name", s)
	}("Todd")
}

func foo() {
	fmt.Println("Foo run!")
}

// a named function with an identifier:
// func (r reciever) indentifier(p parameter(s)) (r return(s)) { code }
// an anonymous function
// func(p parameter(s)) (r return(s)) { code }()
