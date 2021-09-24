package main

func main() {
}

func wiggleMaxLength(nums []int) (ans int) {
	n := len(nums)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i][0] = 1
		dp[i][1] = 1
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				dp[i][0] = max(dp[j][1]+1, dp[i][0])
			}
			if nums[j] < nums[i] {
				dp[i][1] = max(dp[j][0]+1, dp[i][1])
			}
		}
		ans = max(ans, dp[i][0], dp[i][1])
	}
	return ans
}

func max(nums ...int) int {
	ans := nums[0]
	for i := range nums {
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}
