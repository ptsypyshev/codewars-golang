package main

import "fmt"

func OverTheRoad(address int, n int) int {
	return n*2 + 1 - address
}

func main() {
	fmt.Println(OverTheRoad(23633656673, 310027696726), 596421736780)
}
