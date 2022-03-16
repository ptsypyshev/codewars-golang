package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func IsPalindrome(str string) bool {
	lower := strings.ToLower(str)
	lenStr := utf8.RuneCountInString(lower)
	for i := 0; i < lenStr/2; i++ {
		if lower[i] != lower[lenStr-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("Abtab"))
}
