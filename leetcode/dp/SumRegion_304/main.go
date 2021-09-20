package main

import "fmt"

func main() {
	m := Constructor([][]int{{-1, -5}})
	fmt.Println(m.SumRegion(0, 1, 0, 1))
}

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}
	for i := range dp {
		sum := 0
		for j := range dp[i] {
			sum += matrix[i][j]
			if i == 0 {
				dp[i][j] = sum
			} else {
				dp[i][j] = dp[i-1][j] + sum
			}
		}
	}
	return NumMatrix{dp: dp}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	dp := m.dp
	getValue := func(a, b int) int {
		if a < 0 || b < 0 {
			return 0
		}
		if a >= len(dp) {
			a = len(dp) - 1
		}
		if b >= len(dp[a]) {
			b = len(dp[a]) - 1
		}
		return dp[a][b]
	}
	return getValue(row2, col2) + getValue(row1-1, col1-1) - getValue(row1-1, col2) - getValue(row2, col1-1)
}
