package main

import "fmt"

func main() {
	foo()
	bar("Alex")

	s := aloha("Nico")
	fmt.Println(s)

	s1, n := dogYears("Todd", 40)
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is ", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

// func (r reciver) idetifier(p prarmeters(s) (return(s)) {<code here>}
