package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("Key = %v, Value = %v; \n", k, v)
	}

	age1 := m["James"]
	fmt.Println("The age of Bond", age1)
	if v, ok := m["James"]; ok {
		fmt.Println("There is a Bond lookup, and here Bond's age is ", v)
	}

	age2 := m["Q"]
	fmt.Println("The age of Q", age2)
	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int", v)
	}
}
