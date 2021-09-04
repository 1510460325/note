package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		nextPermutation(nums)
		fmt.Println(nums)
	}
}

/**
1 2 3 4 5 6 7 8 9
1 2 3 4 5 6 7 9 8
1 2 3 4 5 6 8 7 9
1 2 3 4 5 7 8 9 7
1 2 3 6 7 8 9
*/
func nextPermutation(nums []int) {
	length := len(nums)
	for i := length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := length - 1; j > i; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					sort.Ints(nums[i+1:])
					return
				}
			}
		}
	}
	sort.Ints(nums)
}
