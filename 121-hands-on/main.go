package main

import "fmt"

func main() {
	mx := make(map[string][]string)
	mx["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	mx["moneypenny_jenny"] = []string{"intellegence", "literature", "computer_science"}
	mx["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	mx["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	fmt.Printf("jb = %#v\n", mx)
	for i, v := range mx {
		for n, m := range v {
			fmt.Printf("index:%v\t value: %v\t feature_ind: %v\t features: %v\n", i, v, n, m)
		}
	}
}
