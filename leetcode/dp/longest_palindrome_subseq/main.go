package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseqV2("cbbdc"))
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

func max(nums ...int) int {
	tem := nums[0]
	for _, i := range nums {
		if tem < i {
			tem = i
		}
	}
	return tem
}

func longestPalindromeSubseqV2(s string) int {
	length := len(s)
	if length < 2 {
		return length
	}
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
		dp[i][i] = 1
	}
	value := func(i, j int) int {
		if i > j {
			return 0
		}
		return dp[i][j]
	}
	borderValue := func(i, j int) int {
		if s[i] == s[j] {
			return 2
		}
		return 0
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max(value(i+1, j), value(i, j-1), value(i+1, j-1)+borderValue(i, j))
		}
	}
	return dp[0][length-1]
}
