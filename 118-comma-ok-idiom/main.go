package main

import "fmt"

func main() {
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	delete(an, "George")

	// in go we limit the scope for v like this:
	if v, ok := an["George"]; !ok {
		fmt.Println("Key didn't exist")
	} else {
		fmt.Println("the value prints: ", v)
	}

	// don't limit the scope for t like this:
	t, ok := an["Lucas"]
	if ok {
		fmt.Println("the value prints: ", t)
	} else {
		fmt.Println("Key didn't exist")
	}

	for k, v := range an {
		fmt.Printf("map: k = %v, v = %v\n", k, v)
	}
}
