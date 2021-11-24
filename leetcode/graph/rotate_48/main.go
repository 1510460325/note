package main

import "fmt"

func main() {
	a := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(a)
	fmt.Println(a)
}

func rotate(matrix [][]int) {
	xx, yy := [2]int{0, 0}, [2]int{len(matrix) - 1, len(matrix) - 1}
	for xx[0] < yy[0] {
		tem := make([]int, 0, yy[0]-xx[0]+1)
		for i := yy[0]; i >= xx[0]; i-- {
			tem = append(tem, matrix[i][xx[1]])
		}
		for i, j := xx[0], xx[1]; i <= yy[0]; i++ {
			matrix[i][xx[1]] = matrix[yy[0]][j]
			j++
		}
		for i, j := yy[0], xx[1]; j <= yy[1]; i-- {
			matrix[yy[0]][j] = matrix[i][yy[1]]
			j++
		}
		for i, j := yy[0], yy[1]; i >= xx[0]; i-- {
			matrix[i][yy[1]] = matrix[xx[0]][j]
			j--
		}
		copy(matrix[xx[0]][xx[1]:yy[1]+1], tem)
		xx[0]++
		xx[1]++
		yy[0]--
		yy[1]--
	}
}
