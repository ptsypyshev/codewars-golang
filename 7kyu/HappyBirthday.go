package main

import (
	"fmt"
	"sort"
)

func WrapPresent(height, width, length int) int {
	sorted := []int{height, width, length}
	sort.Ints(sorted)

	return 4*sorted[0] + 2*(sorted[1]+sorted[2]) + 20
}

func main() {
	fmt.Println(WrapPresent(17, 32, 11))
	fmt.Println(WrapPresent(13, 13, 13))
}
