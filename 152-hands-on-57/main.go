package main

import "fmt"

func main() {
	x := sum([]int{11, 12, 13, 14, 15})
	fmt.Println(x)
}

func sum(ii []int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}

// func sum(ii []int) (total int) {
// 	total = 0
// 	for _, v := range ii {
// 		total += v
// 	}
// 	return
// }
