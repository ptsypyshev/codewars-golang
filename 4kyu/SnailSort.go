package main

import "fmt"

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

func changeDirection(d int) int {
	return (d + 1) % 4
}

func Snail(snaipMap [][]int) []int {
	dir := RIGHT
	lenX := len(snaipMap[0]) - 1
	lenY := len(snaipMap) - 1
	x, y, minX, minY := 0, 0, 0, 0
	var result = make([]int, 0, lenX*lenY)

	for {
		if len(result) == len(snaipMap[0])*len(snaipMap) {
			return result
		}

		result = append(result, snaipMap[y][x])

		switch dir {
		case RIGHT:
			if x == lenX {
				dir = changeDirection(dir)
				y += 1
				lenX -= 1
			} else {
				x += 1

			}
		case DOWN:
			if y == lenY {
				dir = changeDirection(dir)
				x -= 1
				lenY -= 1
			} else {
				y += 1
			}
		case LEFT:
			if x == minX {
				dir = changeDirection(dir)
				y -= 1
				minX += 1
			} else {
				x -= 1
			}
		case UP:
			if y-1 == minY {
				dir = changeDirection(dir)
				x += 1
				minY += 1
			} else {
				y -= 1
			}
		default:
			fmt.Println("OUT_OF_RANGE", dir)
			return result
		}
	}
}

func main() {
	//snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	snailMap := [][]int{{1, 2, 3, 10}, {4, 5, 6, 11}, {7, 8, 9, 12}}
	//snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	fmt.Println("snailmap is ", Snail(snailMap))
}
