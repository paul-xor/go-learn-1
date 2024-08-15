package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happend: ", err)
		// log.Println("err happend: ", err)
		// log.Fatalln("err happend: ", err)
		// log.Panicln("err happend: ", err)
		panic(err)
	}
}
