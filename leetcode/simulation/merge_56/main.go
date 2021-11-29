package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{
		{1, 4},
		{4, 5},
	}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	return merge(append(intervals, newInterval))
}

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, item := range intervals {
		if len(ans) == 0 {
			ans = append(ans, item)
			continue
		}
		pre := ans[len(ans)-1]
		if pre[1] < item[0] {
			ans = append(ans, item)
			continue
		}
		pre[1] = max(pre[1], item[1])
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
