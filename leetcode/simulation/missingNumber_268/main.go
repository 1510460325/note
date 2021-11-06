package main

func main() {

}

func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return total - sum
}