package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3,2,1,0,4}))
}

func canJump(nums []int) bool {
	maxIndex := 0
	for i := range nums {
		if maxIndex >= i {
			maxIndex = max(maxIndex, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
