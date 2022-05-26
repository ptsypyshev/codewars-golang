package main

import "fmt"

func FindOdd(seq []int) int {
	var intSet = make(map[int]int)
	for _, v := range seq {
		intSet[v] += 1
	}

	for k, v := range intSet {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
}
