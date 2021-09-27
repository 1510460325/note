package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(13))
}

func numSquares(n int) int {
	coins := make([]int, int(math.Sqrt(float64(n))))
	for i := range coins {
		coins[i] = (i + 1) * (i + 1)
	}
	dp := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for _, c := range coins {
			if c > i {
				continue
			}
			dp[i] = min(dp[i], dp[i-c] + 1)
		}
	}
	return dp[n]
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
