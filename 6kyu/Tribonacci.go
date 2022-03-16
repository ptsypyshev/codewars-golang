package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	res := make([]float64, 0, n)
	res = signature[:]
	for i := 0; i < n-3; i++ {
		res = append(res, res[i]+res[i+1]+res[i+2])
	}
	return res[:n]
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10))
	fmt.Println(Tribonacci([3]float64{0, 0, 1}, 10))
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 0))
}
