package main

import "fmt"

type Person struct {
	first string
	age   int
}

func main() {
	p := Person{
		first: "James",
		age:   42,
	}

	p.speak()
}

func (p Person) speak() {
	fmt.Printf("I am a person: %v and my age is %v\n", p.first, p.age)
}
