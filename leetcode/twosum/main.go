package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{3,2,4}, 6))
}

func twoSum(nums []int, target int) []int {
	ans := []int{-1, -1}
	cp := make([]int, len(nums))
	copy(cp, nums)
	sort.Ints(cp)
	i, j := 0, len(nums)-1
	for {
		if cp[i]+cp[j] > target {
			j--
		} else if cp[i]+cp[j] < target {
			i++
		} else {
			break
		}
	}
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx] == cp[i] && ans[0] == -1 {
			ans[0] = idx
		} else if nums[idx] == cp[j] && ans[1] == -1 {
			ans[1] = idx
		} else {
			continue
		}
		if ans[1] != -1 && ans[0] != -1 {
			return ans
		}
	}
	return ans
}
