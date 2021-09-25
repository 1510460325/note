package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abc", "def"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	value := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return dp[i][j]
	}
	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[i][j] = value(i-1, j-1) + 1
			}
			dp[i][j] = max(dp[i][j], value(i-1, j), value(i, j-1))
		}
	}
	return dp[m-1][n-1]
}

func max(nums ...int) int {
	ans := nums[0]
	for _, i := range nums {
		if i > ans {
			ans = i
		}
	}
	return ans
}
