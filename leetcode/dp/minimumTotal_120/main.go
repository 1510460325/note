package main

import "fmt"

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}

func minimumTotal(triangle [][]int) int {
	const inf = 1<<31 - 1
	dp := make([]int, len(triangle[0]))
	copy(dp, triangle[0])
	for i := 1; i < len(triangle); i++ {
		dpTem := make([]int, len(triangle[i]))
		for j := range triangle[i] {
			tem := inf
			if j < len(dp) {
				tem = min(tem, dp[j])
			}
			if j > 0 {
				tem = min(tem, dp[j-1])
			}
			dpTem[j] = tem + triangle[i][j]
		}
		dp = dpTem
	}
	return min(dp...)
}

func min(nums ...int) (ans int) {
	ans = nums[0]
	for i := range nums {
		if ans > nums[i] {
			ans = nums[i]
		}
	}
	return ans
}
