package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) (result string) {
	wordLower := strings.ToLower(word)
	for _, v := range wordLower {
		if strings.Count(wordLower, string(v)) > 1 {
			result += ")"
		} else {
			result += "("
		}
	}
	return
}

func main() {
	fmt.Println(DuplicateEncode("recede"))
	fmt.Println(DuplicateEncode("(( @"))
	fmt.Println(DuplicateEncode("Success"))
}
