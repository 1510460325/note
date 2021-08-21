package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}

func threeSumClosest(nums []int, target int) (ans int) {

	ans = 1 << 31
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		sum := target - nums[i]
		for j < k {
			if abs(ans-target) > abs(nums[j]+nums[k]+nums[i]-target) {
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

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		target := nums[i] * -1
		for j < k {
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				for j < len(nums) && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k > i && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
