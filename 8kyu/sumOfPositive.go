package main

import "fmt"

func PositiveSum(numbers []int) int {
	var res int
	for _, v := range numbers {
		if v > 0 {
			res += v
		}
	}
	return res
}

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}))
	fmt.Println(PositiveSum([]int{-1, -2, -3, -4, -5}))
	fmt.Println(PositiveSum([]int{}))
	fmt.Println(PositiveSum([]int{-1, 2, 3, 4, -5}))
}
