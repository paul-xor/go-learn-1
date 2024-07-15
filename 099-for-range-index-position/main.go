package main

import "fmt"

func main() {
	sl := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Printf("slice is %v\n %v\n", sl, len(sl))

	// blank identifier to not use a returned value
	for _, v := range sl {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("-------")
	fmt.Println(sl[0])
	fmt.Println(sl[1])
	fmt.Println(sl[2])
}
