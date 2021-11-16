package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{1, 2, 3, 5, 6, 7}, 9))
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(idx, val int) {
		if idx == len(candidates) {
			return
		}
		if val == target {
			newAns := make([]int, len(path))
			copy(newAns, path)
			ans = append(ans, newAns)
			return
		} else if val > target {
			return
		} else if val+candidates[idx] <= target {
			path = append(path, candidates[idx])
			dfs(idx, val+candidates[idx])
			path = path[:len(path)-1]
			dfs(idx+1, val)
		}
	}
	path = append(path, candidates[0])
	dfs(0, candidates[0])
	path = make([]int, 0)
	dfs(1, 0)
	return ans
}
