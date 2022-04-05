package main

import (
	"fmt"
	"reflect"
)

func getCharMap(word string) map[string]int {
	res := make(map[string]int)
	for _, elem := range word {
		res[string(elem)] += 1
	}
	return res
}

func Anagrams(word string, words []string) (res []string) {
	orig := getCharMap(word)
	for _, elem := range words {
		if reflect.DeepEqual(orig, getCharMap(elem)) {
			res = append(res, elem)
		}
	}
	return
}

func main() {
	fmt.Println(Anagrams("abba", []string{"aabb", "abcd", "bbaa", "dada"}))
	fmt.Println(Anagrams("laser", []string{"lazing", "lazy", "lacer"}))
}
