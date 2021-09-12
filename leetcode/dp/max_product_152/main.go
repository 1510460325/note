package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) (ans int) {
	minVal, maxVal := nums[0], nums[0]
	ans = maxVal
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxVal, minVal = max(minVal*nums[i], nums[i]), min(maxVal*nums[i], nums[i])
		} else {
			maxVal, minVal = max(maxVal*nums[i], nums[i]), min(minVal*nums[i], nums[i])
		}
		ans = max(maxVal, ans)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
