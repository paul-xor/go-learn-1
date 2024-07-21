package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func saySomethig(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	p2 := person{
		first: "Jenny",
	}
	sa1.speak()
	p2.speak()

	// another way to call speak()
	saySomethig(sa1)
	saySomethig(p2)

}

// func (r reciever) identifier(p paremeter(s)) (r return(s)) {code}

func (p person) speak() {
	fmt.Println("I am ", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}
