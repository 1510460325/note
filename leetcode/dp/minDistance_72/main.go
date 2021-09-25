package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n - m
	} else if n == 0 {
		return m - n
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range word1 {
		if word1[i] == word2[0] {
			dp[i][0] = i
		} else if i > 0 {
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = i + 1
		}
	}
	for i := range word2 {
		if word2[i] == word1[0] {
			dp[0][i] = i
		} else if i > 0 {
			dp[0][i] = dp[0][i-1] + 1
		} else {
			dp[i][0] = i + 1
		}
	}
	const inf = 1<<31 - 1
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = inf
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
		}
	}
	return dp[m-1][n-1]
}

func min(nums ...int) int {
	ans := nums[0]
	for _, i := range nums {
		if ans > i {
			ans = i
		}
	}
	return ans
}
