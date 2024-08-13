package main

import "fmt"

func main() {
	// this works with bufferd chan or with commented goroutine
	c := make(chan int, 1)

	// go func() {
	// 	c <- 42
	// }()
	c <- 42

	fmt.Println(<-c)
}
