package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) (ans int) {

	ans = 1 << 31
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		sum := target - nums[i]
		for j < k {
			if abs(ans-target) > abs(nums[j]+nums[k]+nums[i] - target) {
				ans = nums[j] + nums[k] + nums[i]
			}
			if nums[j]+nums[k] > sum {
				k--
			} else if nums[j]+nums[k] < sum {
				j++
			} else {
				return target
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
