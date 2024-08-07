package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[0]
	for index, arg := range os.Args[1:] {
		fmt.Printf("# %v, arg: %v\n", index, arg)
	}
	fmt.Println(s)
}
