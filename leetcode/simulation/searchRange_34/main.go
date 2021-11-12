package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1}, 1))
}

func searchRange(nums []int, target int) []int {
	return []int{findTarget(nums, target, true),findTarget(nums, target, false)}
}

func findTarget(nums []int, target int, left bool) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else if left {
			if m > 0 && nums[m-1] == target {
				r = m
			} else {
				return m
			}
		} else {
			if m < len(nums)-1 && nums[m+1] == target {
				l = m + 1
			} else {
				return m
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
