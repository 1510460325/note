package main

import "fmt"

func main() {
	fmt.Println(getMaximumGenerated(11))
}

func getMaximumGenerated(n int) (ans int) {
	if n == 0 {
		return 0
	}
	nums := make([]int, n+1)
	nums[1] = 1
	ans = 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}
