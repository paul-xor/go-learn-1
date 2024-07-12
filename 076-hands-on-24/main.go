package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)

	fmt.Printf("The random value: %v \n", x)
	if x <= 100 {
		fmt.Println("the value is between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("the value is between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("the value is between 101 and 250")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(3))
	}

}
