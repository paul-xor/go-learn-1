package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favIce    []string
}

func main() {
	person1 := person{
		firstName: "James",
		lastName:  "Bond",
		favIce:    []string{"chocolate", "creambrule"},
	}

	person2 := person{
		firstName: "Jenny",
		lastName:  "Moneypenny",
		favIce:    []string{"vanilla", "strawberry", "mint chocolate chip"},
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println("-----------")
	fmt.Printf("Person#1 - %v, %v likes: \n", person1.firstName, person1.lastName)
	for i, v := range person1.favIce {
		fmt.Println(i, v)
	}

	fmt.Printf("Person#2 - %v, %v likes: \n", person2.firstName, person2.lastName)
	for i, v := range person2.favIce {
		fmt.Println(i, v)
	}

}
