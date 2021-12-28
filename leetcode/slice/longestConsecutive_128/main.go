package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
}

func longestConsecutive(nums []int) (ret int) {
	sort.Ints(nums)
	cur, pre := 0, 0
	for i := range nums {
		if i > 0 {
			if pre == nums[i] {
				continue
			}
			if nums[i] != pre+1 {
				ret = max(ret, cur)
				cur = 0
			}
		}
		cur++
		pre = nums[i]
	}
	return max(ret, cur)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
