package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{9}, 3))
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < target+1; i++ {
		for _, j := range nums {
			if j <= i {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[target]
}
