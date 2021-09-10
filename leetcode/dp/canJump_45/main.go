package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	const inf = 1<<31 - 1
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}
	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
