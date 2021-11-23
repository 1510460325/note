package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for {
		now := make([]int, len(nums))
		copy(now, nums)
		ans = append(ans, now)
		if !next(nums) {
			return ans
		}
	}
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
