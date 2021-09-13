package main

import "fmt"

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{1, 2}))
}

func maxScoreSightseeingPair(values []int) int {
	ans, mx := 0, values[0]+0
	for j := 1; j < len(values); j++ {
		ans = max(ans, mx+values[j]-j)
		mx = max(mx, values[j]+j)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
