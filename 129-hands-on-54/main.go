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

	m := map[string]person{
		person1.last_name: person1,
		person2.last_name: person2,
	}

	fmt.Printf("-> First person: %v\t %v\n", m["Bond"].first_name, m["Bond"].last_name)
	for i, v := range m["Bond"].fav_ice {
		fmt.Println(i, v)
	}

	fmt.Printf("-> Second preson: %v\t %v\n", m["Moneypenny"].first_name, m["Moneypenny"].last_name)
	for i, v := range m["Moneypenny"].fav_ice {
		fmt.Println(i, v)
	}

	// or we range like this:
	fmt.Println("---------")
	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.fav_ice {
			fmt.Printf("name: %v\t %v\t %v\n", v.first_name, v.last_name, v2)
		}
	}

}
