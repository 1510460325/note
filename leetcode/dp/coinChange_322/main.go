package main

import (
	"fmt"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	const inf = 1<<31 - 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] == inf {
		return -1
	}
	return dp[amount]
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
