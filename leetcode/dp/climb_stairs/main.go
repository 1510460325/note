package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n)
	dp[0], dp[1] = 1, 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func minCostClimbingStairs(cost []int) int {
	height := len(cost)
	if height < 2 {
		return 0
	}
	dp := make([]int, height+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= height; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[height]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
