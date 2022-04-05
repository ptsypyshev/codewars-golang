package main

import "fmt"

func unpackList(list []int, res string) string {
	if len(list) > 2 {
		res += fmt.Sprintf("%d-%d,", list[0], list[len(list)-1])
	} else {
		for _, elem := range list {
			res += fmt.Sprintf("%d,", elem)
		}
	}
	return res
}

func Solution(list []int) (res string) {
	tmpList := make([]int, 0, 10)
	for idx, val := range list {
		if idx == 0 {
			tmpList = append(tmpList, val)
		} else {
			if len(tmpList) > 0 && tmpList[len(tmpList)-1]+1 == val {
				tmpList = append(tmpList, val)
			} else {
				res = unpackList(tmpList, res)
				tmpList = make([]int, 1, 10)
				tmpList[0] = val
			}
		}

	}
	res = unpackList(tmpList, res)
	return res[:len(res)-1]
}

func main() {
	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
}
