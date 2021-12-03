package main

import "fmt"

func main() {
	fmt.Println(combine(5, 4))
}

func combine(n int, k int) (ans [][]int) {
	var dfs func(path []int, idx, count int)
	dfs = func(path []int, idx, rest int) {
		if rest == 0 {
			newAns := make([]int, len(path))
			copy(newAns, path)
			ans = append(ans, newAns)
			return
		}
		if idx > n {
			return
		}
		dfs(append(path, idx), idx+1, rest-1)
		dfs(path, idx+1, rest)
	}
	dfs(nil, 1, k)
	return ans
}
