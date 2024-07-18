package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	p2 := person{
		first: "Jenny",
		last:  "MoneyPenny",
		age:   32,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v \n", p1, p1)
	fmt.Printf("His name: %v %v, age: %v\n", p1.first, p1.last, p1.age)

}
