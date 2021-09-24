package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{1,3,6,7,9,4,10,5,6}))
}

func lengthOfLIS(nums []int) (ans int) {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}
