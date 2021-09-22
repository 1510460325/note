package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	const inf = 1<<31 - 1
	for i := range dp {
		for j := range dp[i] {
			tem := inf
			if i > 0 {
				tem = min(tem, dp[i-1][j])
			}
			if j > 0 {
				tem = min(tem, dp[i][j-1])
			}
			if tem == inf {
				dp[i][j] = grid[i][j]
			} else {
				dp[i][j] = grid[i][j] + tem
			}
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
