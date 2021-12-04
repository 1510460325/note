package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) (ans [][]int) {
	var dfs func(path []int, idx, rest int)
	dfs = func(path []int, idx, rest int) {
		if rest == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		if idx >= len(nums) {
			return
		}
		dfs(append(path, nums[idx]), idx+1, rest-1)
		dfs(path, idx+1, rest)
	}
	for i := 0; i <= len(nums); i++ {
		dfs(nil, 0, i)
	}
	return ans
}
