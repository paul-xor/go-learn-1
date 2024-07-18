package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	fav_ice    []string
}

func main() {
	person1 := person{
		first_name: "James",
		last_name:  "Bond",
		fav_ice:    []string{"chocolate", "creambrule"},
	}

	person2 := person{
		first_name: "Jenny",
		last_name:  "Moneypenny",
		fav_ice:    []string{"vanilla", "strawberry", "mint chocolate chip"},
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println("-----------")
	fmt.Printf("Person#1 - %v, %v likes: \n", person1.first_name, person1.last_name)
	for i, v := range person1.fav_ice {
		fmt.Println(i, v)
	}

	fmt.Printf("Person#2 - %v, %v likes: \n", person2.first_name, person2.last_name)
	for i, v := range person2.fav_ice {
		fmt.Println(i, v)
	}

}
