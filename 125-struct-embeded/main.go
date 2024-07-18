package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk   bool
	first string
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   42,
		},
		ltk:   true,
		first: "ETHAN HAWK",
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)
	fmt.Println(sa1.first, sa1.ltk, sa1.person.first, sa1.person.last)
}
