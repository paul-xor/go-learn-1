package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

// wrapper function
func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 140", s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}
	var n count = 42

	log.Println(b)
	log.Println(n)
	fmt.Println("--------")
	logInfo(b)
	logInfo(n)
}
