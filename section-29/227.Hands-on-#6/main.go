package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		v, ok := <-c
		if !ok {
			fmt.Println("All values received, and channel closed.")
			break
		}
		fmt.Println(v)
	}

}
