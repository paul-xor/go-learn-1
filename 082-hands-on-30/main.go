package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		switchState(x, i)
	}
}

func switchState(x, i int) {
	switch x {
	case 0:
		fmt.Printf("iteration# = %v, x is 0; \n", i)
	case 1:
		fmt.Printf("iteration# = %v, x is 1; \n", i)
	case 2:
		fmt.Printf("iteration# = %v, x is 2; \n", i)
	case 3:
		fmt.Printf("iteration# = %v, x is 3; \n", i)
	case 4:
		fmt.Printf("iteration# = %v, x is 4; \n", i)
	}
}
