package main

func main() {

}

func runningSum(nums []int) []int {
	sum := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			sum[i] = nums[i]
		} else {
			sum[i] = nums[i] + sum[i-1]
		}
	}
	return sum
}
