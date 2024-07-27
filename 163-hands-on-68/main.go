package main

import "fmt"

func main() {
	func() {
		fmt.Println("this is anonymous func")
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
}
