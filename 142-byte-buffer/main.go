package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers!")
	fmt.Println(b.String())
	b.Reset()
	fmt.Println(b.String())
	b.WriteString("Hello again, it is Wednesday.")
	fmt.Println(b.String())
	b.Write([]byte(" ... Happy, happy"))
	fmt.Println(b.String())
}
