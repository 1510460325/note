package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
}

func fourSum(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			l, r := j+1, len(nums)-1
			towSum := target - nums[i] - nums[j]
			for l < r {
				if nums[l]+nums[r] == towSum {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < len(nums) && nums[l] == nums[l-1] {
						l++
					}
					for r > j && nums[r] == nums[r+1] {
						r--
					}
				} else if nums[l]+nums[r] < towSum {
					l++
				} else {
					r--
				}
			}
			j++
			for j < len(nums)-2 && nums[j] == nums[j-1] {
				j++
			}
		}
		i++
		for i < len(nums)-3 && nums[i] == nums[i-1] {
			i++
		}
	}
	return ans
}
