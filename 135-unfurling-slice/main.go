package main

import "fmt"

func main() {
	xi := []int{1, 3, 5, 6, 7}
	s := sum(xi...)
	fmt.Println("The sum is ", s)
}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

// func (r reciver) idetifier(p prarmeters(s) (return(s)) {<code here>}
