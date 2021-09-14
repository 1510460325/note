package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := [][]int{
		{0, -prices[0]},
		{max(0, prices[1]-prices[0]), max(-prices[0], -prices[1])},
	}
	for i := 2; i < len(prices); i++ {
		// dp0 = max(前两天不持股，前一天不持股，当前刚卖)
		dp0 := max(dp[0][0], dp[1][0], dp[1][1]+prices[i])
		// dp1 = max(前两天持股，前一天持股，当前刚买)
		dp1 := max(dp[0][1], dp[1][1], dp[0][0]-prices[i])
		dp[0][0], dp[0][1], dp[1][0], dp[1][1] = dp[1][0], dp[1][1], dp0, dp1
	}
	return dp[1][0]
}

func max(a ...int) int {
	m := a[0]
	for i := range a {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}
