package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, -1}))
}

func maxSubArray(nums []int) int {
	sum, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(sum+nums[i], nums[i])
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
