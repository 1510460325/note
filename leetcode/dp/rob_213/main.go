package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = nums[0]
	dp[1][0] = nums[1]
	dp[1][1] = nums[0]
	for i := 2; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-2][0]+nums[i])
		if i == length-1 {
			dp[i][1] = dp[i-1][1]
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-2][1]+nums[i])
		}
	}
	return max(dp[length-1][0], dp[length-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
