package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is an error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1)
}

func foo(e error) {
	// here is assertion:
	fmt.Println("foo run - ", e, "\n", e.(customErr).info)
}
