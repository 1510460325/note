package main

import "fmt"

func main() {
	fmt.Println(thirdMax([]int{-2147483648, 1, 2}))
}

func thirdMax(nums []int) int {
	rank := [3]int{-1 << 31, -1 << 31, -1 << 31}
	top := 0
	for _, i := range nums {
		if rank[0] <= i {
			if rank[0] < i {
				rank[0], rank[1], rank[2] = i, rank[0], rank[1]
				top++
			}
			if top < 1 {
				top++
			}
		} else if rank[1] <= i {
			if rank[1] < i {
				rank[1], rank[2] = i, rank[1]
				top++
			}
			if top < 2 {
				top = 2
			}
		} else if rank[2] <= i {
			rank[2] = i
			top = 3
		}
	}
	if top >= 3 {
		return rank[2]
	}
	return rank[0]
}
