package main

import (
	"fmt"
)

func main() {
	fmt.Println(minMoves([]int{1, 1, 10000000}))
}

func minMoves(nums []int) (ans int) {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		ans += num - min
	}
	return
}
