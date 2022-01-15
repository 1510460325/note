package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{6, 3, 5, 2, 4}))
}

func findPeakElement(nums []int) (rst int) {

	rst = -1
	var dfs func(s, e int)

	dfs = func(s, e int) {
		if rst != -1 || s > e {
			return
		}
		fmt.Println(s, e)
		if s == e || s == e-1 {
			if nums[s] < nums[e] {
				rst = e
			} else {
				rst = s
			}
			return
		}
		mid := (s + e) / 2
		if nums[mid] > nums[mid+1] {
			dfs(s, mid)
		} else {
			dfs(mid+1, e)
		}
	}
	dfs(0, len(nums)-1)
	return rst
}
