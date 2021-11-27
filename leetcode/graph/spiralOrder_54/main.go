package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{
		{7},
		{9},
		{6},
	}))
}

type location struct {
	x, y int
}

func spiralOrder(matrix [][]int) (ret []int) {
	a, b := location{0, 0}, location{len(matrix) - 1, len(matrix[0]) - 1}
	for a.x <= b.x && a.y <= b.y {
		for j := a.y; j <= b.y; j++ {
			ret = append(ret, matrix[a.x][j])
		}
		if a.x == b.x {
			break
		}
		for i := a.x + 1; i <= b.x; i++ {
			ret = append(ret, matrix[i][b.y])
		}
		for j := b.y - 1; j >= a.y; j-- {
			ret = append(ret, matrix[b.x][j])
		}
		if a.y == b.y {
			break
		}
		for i := b.x - 1; i > a.x; i-- {
			ret = append(ret, matrix[i][a.y])
		}
		a.x++
		a.y++
		b.x--
		b.y--
	}
	return ret
}
