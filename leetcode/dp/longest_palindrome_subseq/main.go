package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("aaab"))
}

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for step := 1; step <= len(s); step++ {
		for i := 0; i < len(s); i++ {
			j := i + step
			if j >= len(s) {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			} else if step == 1 {
				dp[i][j] = 2
			} else {
				dp[i][j] = dp[i+1][j-1] + 2
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
