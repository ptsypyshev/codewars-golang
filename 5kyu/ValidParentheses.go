package main

import "fmt"

func ValidParentheses(parens string) bool {
	var validCounter int
	for _, elem := range parens {
		if elem == '(' {
			validCounter++
		}
		if elem == ')' {
			validCounter--
		}
		if validCounter < 0 {
			return false
		}
	}
	if validCounter != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(ValidParentheses("(())((()())())"))
}
