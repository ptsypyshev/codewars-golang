package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NbDig(n int, d int) (res int) {
	dStr := strconv.Itoa(d)
	for i := 0; i <= n; i++ {
		resultStr := strconv.Itoa(i * i)
		res += strings.Count(resultStr, dStr)
	}
	return
}

func main() {
	fmt.Println(NbDig(550, 5))
	fmt.Println(NbDig(5750, 0))
}
