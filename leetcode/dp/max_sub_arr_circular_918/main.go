package main

import "fmt"

func main() {
	fmt.Println(maxSubarraySumCircular([]int{-2}))
}

func maxSubarraySumCircular(num []int) int {
	maxSum := num[0]
	maxAns := maxSum
	for i := 1; i < len(num); i++ {
		maxSum = max(maxSum+num[i], num[i])
		maxAns = max(maxAns, maxSum)
	}
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	minSum := 0
	minAns := 0
	for i := 1; i < len(num)-1; i++ {
		minSum = min(minSum+num[i], num[i])
		minAns = min(minSum, minAns)
	}
	return max(maxAns, sum-minAns)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
