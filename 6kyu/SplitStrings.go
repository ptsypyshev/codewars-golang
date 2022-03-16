package main

import "fmt"

func Solution(str string) (res []string) {
	var tmp string
	for i, v := range str {
		tmp += string(v)
		if i%2 != 0 {
			res = append(res, tmp)
			tmp = ""
		}
	}
	if len(tmp) != 0 {
		res = append(res, tmp+"_")
	}
	return
}

func main() {
	fmt.Println(Solution("awsaws"))
	fmt.Println(Solution("dawsd"))
}
