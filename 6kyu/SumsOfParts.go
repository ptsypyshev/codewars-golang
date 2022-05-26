package main

import "fmt"

func sumNumbersInSlice(inSlice []uint64) (result uint64) {
	for _, v := range inSlice {
		result += v
	}
	return
}

func PartsSums(ls []uint64) []uint64 {
	var result = make([]uint64, 0, len(ls)+1)
	result = append(result, sumNumbersInSlice(ls))
	for _, v := range ls {
		result = append(result, result[len(result)-1]-v)
	}
	return result
}

func main() {
	fmt.Println(PartsSums([]uint64{1, 2, 3, 4, 5, 6}), []uint64{21, 20, 18, 15, 11, 6, 0})
}
