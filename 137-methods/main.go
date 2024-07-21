package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "James",
	}
	p2 := person{
		first: "Jenny",
	}
	p1.speak()
	p2.speak()

}

// func (r reciever) identifier(p paremeter(s)) (r return(s)) {code}

func (p person) speak() {
	fmt.Println("I am ", p.first)
}
