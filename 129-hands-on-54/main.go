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

	m := map[string]person{
		person1.lastName: person1,
		person2.lastName: person2,
	}

	fmt.Printf("-> First person: %v\t %v\n", m["Bond"].firstName, m["Bond"].lastName)
	for i, v := range m["Bond"].favIce {
		fmt.Println(i, v)
	}

	fmt.Printf("-> Second preson: %v\t %v\n", m["Moneypenny"].firstName, m["Moneypenny"].lastName)
	for i, v := range m["Moneypenny"].favIce {
		fmt.Println(i, v)
	}

	// or we range like this:
	fmt.Println("---------")
	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIce {
			fmt.Printf("name: %v\t %v\t %v\n", v.firstName, v.lastName, v2)
		}
	}

}
