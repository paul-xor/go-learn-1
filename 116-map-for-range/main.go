package main

import "fmt"

func main() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	for k, v := range an {
		fmt.Printf("map: k = %v, v = %v\n", k, v)
	}
}
