package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	var res []string
	splitted := strings.Split(text, " ")
	fmt.Println(splitted)
	for _, v := range splitted {
		word := []rune(v)
		lenWord := len(word)
		w0 := ""
		if lenWord > 2 {
			word[1], word[lenWord-1] = word[lenWord-1], word[1]
		}
		if lenWord > 0 {
			w0 = strconv.Itoa(int(word[0]))
			res = append(res, w0+string(word[1:]))
		}

	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))
	fmt.Println(EncryptThis(""))
}
