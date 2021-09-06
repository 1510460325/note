package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) (ans int) {
	if len(height) < 3 {
		return 0
	}
	leftMax := make([]int, len(height))
	preMax := height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = preMax
		preMax = max(preMax, height[i])
	}
	rightMax := make([]int, len(height))
	preMax = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = preMax
		preMax = max(preMax, height[i])
	}
	for i := 1; i < len(height)-1; i++ {
		if h := min(leftMax[i], rightMax[i]); h > height[i] {
			ans += h - height[i]
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
