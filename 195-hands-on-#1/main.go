package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   42,
	}
	p2 := person{
		First: "Jenny",
		Last:  "Moneypenny",
		Age:   25,
	}
	p3 := person{
		First: "Dr No",
		Last:  "None",
		Age:   37,
	}

	people := []person{p1, p2, p3}

	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(bs)
}
