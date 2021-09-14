package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]-fee
	for i := 1; i < len(prices); i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i]-fee)
	}
	return dp0
}

func max(a ...int) int {
	m := a[0]
	for i := range a {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}
