package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	numbers := strings.Split(in, " ")
	var high, low int
	for k, v := range numbers {
		res, err := strconv.Atoi(v)
		if err != nil {
			return "Error"
		}
		if k == 0 {
			high, low = res, res
			continue
		}
		if res < low {
			low = res
		}
		if res > high {
			high = res
		}
	}
	return fmt.Sprintf("%d %d", high, low)
}

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
