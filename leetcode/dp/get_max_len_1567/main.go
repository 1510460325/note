package main

import "fmt"

func main() {
	fmt.Println(getMaxLen([]int{1,2,3,5,-6,4,0,10}))
}

func getMaxLen(nums []int) (ans int) {
	negativeLen, positiveLen := 0, 0
	for i := range nums {
		if nums[i] < 0 {
			if negativeLen > 0 {
				positiveLen, negativeLen = negativeLen+1, positiveLen+1
			} else {
				positiveLen, negativeLen = 0, positiveLen+1
			}
			ans = max(positiveLen, ans)
		} else if nums[i] > 0 {
			positiveLen++
			if negativeLen > 0 {
				negativeLen++
			}
			ans = max(positiveLen, ans)
		} else {
			positiveLen, negativeLen = 0, 0
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
