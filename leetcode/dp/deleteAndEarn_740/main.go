package main

func main() {

}

func deleteAndEarn(nums []int) int {
	maxValue := -1
	for _, item := range nums {
		maxValue = max(item, maxValue)
	}
	values := make([]int, maxValue+1)
	for _, item := range nums {
		values[item] += item
	}
	if len(values) == 1 {
		return values[0]
	} else if len(values) == 2 {
		return max(values[0], values[1])
	}
	first, second := values[0], max(values[0], values[1])
	for i := 2; i < len(values); i++ {
		first, second = second, max(first+values[i], second)
	}
	return second
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}