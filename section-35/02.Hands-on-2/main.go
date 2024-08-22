package main

import (
	"fmt"
	"whole_learn_module/section-35/02.Hands-on-2/quote"
	"whole_learn_module/section-35/02.Hands-on-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
