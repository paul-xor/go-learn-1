package main

import "fmt"

func main() {
	i := 0
	// this is infinite loop
	for {
		fmt.Printf("loop iteration# %v \n", i)
		if i == 50 {
			break
		}
		i++
	}
}
