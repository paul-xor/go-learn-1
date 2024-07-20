package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jenny": 27,
			"Q":     87,
			"Ian":   47,
		},
		favDrinks: []string{
			"Martini",
			"Wisky",
		},
	}

	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends - ", k, v)
	}

	for _, v := range p1.favDrinks {
		fmt.Println(p1.first, "- drinks - ", v)
	}

}
