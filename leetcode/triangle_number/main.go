package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{2,2,3,4}))
}


func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	l := len(nums)
	for i := 0; i < l - 2; i++ {
		for j := i + 1; j < l - 1; j++ {
			for k := l -1; k > j; k-- {
				if  nums[i] + nums[j] > nums[k] {
					ans += k - j
					break
				}
			}
		}
	}
	return ans
}