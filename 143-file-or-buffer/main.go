package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, error := w.Write([]byte(p.first))
	return error
}

func main() {
	p := person{
		first: "Jenny",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()
	var b bytes.Buffer
	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}
