package main

import "fmt"

func main() {
	fmt.Println(minFallingPathSum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))
}

func minFallingPathSum(matrix [][]int) int {
	dp := make([]int, len(matrix[0]))
	for i := range matrix[0] {
		dp[i] = matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		dpTemp := make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[i]); j++ {
			tem := dp[j]
			if j > 0 {
				tem = min(tem, dp[j-1])
			}
			if j < len(matrix[i])-1 {
				tem = min(tem, dp[j+1])
			}
			dpTemp[j] = tem + matrix[i][j]
		}
		dp = dpTemp
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
