package main

import (
	"fmt"
)

func main() {

	fmt.Println(allPathsSourceTarget([][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {4, 6}, {5, 6, 7}, {6}, {7}, {}}))
}

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	stk := []int{0}
	var dfs func(int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int{}, stk...))
			return
		}
		for _, y := range graph[x] {
			stk = append(stk, y)
			dfs(y)
			stk = stk[:len(stk)-1]
		}
	}
	dfs(0)
	return
}
