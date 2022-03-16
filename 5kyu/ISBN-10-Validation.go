package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	var res int
	for k, v := range isbn {
		vInt, err := strconv.Atoi(string(v))
		if err != nil {
			if k != 9 || strings.ToLower(string(v)) != "x" {
				return false
			}
			vInt = 10
		}
		res += (k + 1) * vInt
	}
	return res%11 == 0
}

func main() {
	fmt.Println(ValidISBN10("1112223339"))
	fmt.Println(ValidISBN10("1293"))
	fmt.Println(ValidISBN10("1234512345"))
	fmt.Println(ValidISBN10("048665088x"))
}
