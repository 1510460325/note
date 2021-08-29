package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{2}))
}

func sumOddLengthSubarrays(arr []int) (ans int) {
	preSum := make([]int, len(arr))
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		preSum[i] = sum
	}
	for length := 1; length <= len(arr); length += 2 {
		ans += preSum[length-1]
		for i := length; i < len(arr); i++ {
			ans += preSum[i] - preSum[i-length]
		}
	}
	return ans
}
