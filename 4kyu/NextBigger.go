package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func reverseString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] > r[j]
	})
	return string(r)
}

func NextBigger(n int) int {
	strN := strconv.Itoa(n)
	if strN == reverseString(strN) {
		return -1
	}
	if max, err := strconv.Atoi(strings.Repeat("9", len(strN))); err != nil {
		return -1
	} else {
		for i := n + 1; i < max; i++ {
			flagTrue := true
			strStep := strconv.Itoa(i)
			for _, elem := range strN {
				if strings.Count(strStep, string(elem)) != strings.Count(strN, string(elem)) {
					flagTrue = false
					break
				}
			}
			if flagTrue {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(NextBigger(414))
}
