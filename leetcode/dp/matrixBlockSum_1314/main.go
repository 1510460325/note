package main

import "fmt"

func main() {
	fmt.Println(matrixBlockSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, 2))
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	dp := make([][]int, len(mat))
	for i := range dp {
		dp[i] = make([]int, len(mat[i]))
	}
	for i := range dp {
		sum := 0
		for j := range dp[i] {
			sum += mat[i][j]
			if i == 0 {
				dp[i][j] = sum
			} else {
				dp[i][j] = dp[i-1][j] + sum
			}
		}
	}
	ans := make([][]int, len(mat))
	for i := range ans {
		ans[i] = make([]int, len(mat[i]))
	}
	getValue := func(a, b int) int {
		if a < 0 || b < 0 {
			return 0
		}
		if a >= len(mat) {
			a = len(mat) - 1
		}
		if b >= len(mat[a]) {
			b = len(mat[a]) - 1
		}
		return dp[a][b]
	}
	for i := range ans {
		for j := range ans[i] {
			ans[i][j] = getValue(i+k, j+k) + getValue(i-k-1, j-k-1) - getValue(i+k, j-k-1) - getValue(i-k-1, j+k)
		}
	}
	return ans
}
