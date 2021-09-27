package main

import "fmt"

func main() {
	fmt.Println(integerBreak(8))
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = 1
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max(j, dp[j])*max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

func max(nums ...int) int {
	ans := nums[0]
	for _, i := range nums {
		if ans < i {
			ans = i
		}
	}
	return ans
}
