package main

import "fmt"

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1))
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 2, 3, 4, 5}, 5))
}

func findPoisonedDuration(timeSeries []int, duration int) (ans int) {
	for i := range timeSeries {
		if i < len(timeSeries) - 1 {
			ans += min(duration, timeSeries[i+1] - timeSeries[i])
		} else {
			ans += duration
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
