package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permute([]int{1}))
}

func permute(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for {
		newAns := make([]int, len(nums))
		copy(newAns, nums)
		ans = append(ans, newAns)
		if !next(nums) {
			break
		}
	}
	return ans
}

func next(nums []int) bool {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					sort.Ints(nums[i+1:])
					return true
				}
			}
		}
	}
	return false
}