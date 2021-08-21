package main

import "fmt"

func main() {
	fmt.Println(findPaths(36, 5, 50, 15, 3))
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) (ans int) {
	if maxMove == 0 {
		return 0
	}
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxMove+1)
		}
	}
	for i := 0; i < m; i++ {
		dp[i][0][1]++
		dp[i][n-1][1]++
	}
	for j := 0; j < n; j++ {
		dp[0][j][1]++
		dp[m-1][j][1]++
	}
	ans = dp[startRow][startColumn][1]
	for move := 2; move <= maxMove; move++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i > 0 {
					dp[i][j][move] += dp[i-1][j][move-1]
					dp[i][j][move] %= 1e9 + 7
				}
				if i < m-1 {
					dp[i][j][move] += dp[i+1][j][move-1]
					dp[i][j][move] %= 1e9 + 7
				}
				if j > 0 {
					dp[i][j][move] += dp[i][j-1][move-1]
					dp[i][j][move] %= 1e9 + 7
				}
				if j < n-1 {
					dp[i][j][move] += dp[i][j+1][move-1]
					dp[i][j][move] %= 1e9 + 7
				}
			}
		}
		ans += dp[startRow][startColumn][move]
		ans %= 1e9 + 7
	}
	return ans
}
