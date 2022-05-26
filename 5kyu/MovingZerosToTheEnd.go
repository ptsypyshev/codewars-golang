package main

import "fmt"

func MoveZeros(arr []int) []int {
	result := make([]int, 0, len(arr))
	var zeroCounter int
	for _, val := range arr {
		if val != 0 {
			result = append(result, val)
		} else {
			zeroCounter++
		}
	}
	zeros := make([]int, zeroCounter)
	result = append(result, zeros...)
	return result
}

func main() {
	fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}), []int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0})
}
