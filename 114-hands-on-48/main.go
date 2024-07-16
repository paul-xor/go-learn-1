package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}

	xxs := [][]string{jb, jm}
	fmt.Printf("xxs = %v\n", xxs)

	for k, v := range xxs {
		fmt.Println("-----")
		fmt.Printf("k - %v,\t v = %v\n", k, v)
		for _, n := range v {
			fmt.Println(n)
		}
	}

}
