package main

import (
	"fmt"
	"sort"
)

const maxLenLinearSet = 1000000

var linearSet = make(map[int]bool)
var keysSlice []int

func prePopulateLinearSlice() {
	linearSet[1] = false

	for len(linearSet) < maxLenLinearSet {
		for k, v := range linearSet {
			if v {
				continue
			}
			linearSet[2*k+1] = false
			linearSet[3*k+1] = false
			linearSet[k] = true
		}
	}
	keysSlice = make([]int, 0, len(linearSet))
	for k := range linearSet {
		keysSlice = append(keysSlice, k)
	}
	sort.Ints(keysSlice)
}

func DblLinear(n int) (res int) {
	if len(linearSet) == 0 {
		prePopulateLinearSlice()
	}
	return keysSlice[n]
}

func main() {
	fmt.Println(DblLinear(10), 22)
	fmt.Println(DblLinear(50), 175)
	fmt.Println(DblLinear(100), 447)
	fmt.Println(DblLinear(500), 3355)
	fmt.Println(DblLinear(1000), 8488)
}
