package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is from the init function!")
}

func main() {
	for i := 0; i < 4; i++ {
		x := rand.Intn(300)
		fmt.Printf("the init value is: %v \n", x)
		switchFunc(x)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%v \t", rand.Intn(3))
	}

}

func switchFunc(x int) {
	switch {
	case x <= 100:
		fmt.Println("the value is between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("the value is between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("the value is between 201 and 250")
	default:
		fmt.Println("this was more than 250")
	}
	fmt.Println("<===end of switch===>")
}
