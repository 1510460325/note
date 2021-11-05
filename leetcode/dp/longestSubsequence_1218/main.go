package main

func main() {

}

func longestSubsequence(arr []int, difference int) (ans int) {
	endCount := make(map[int]int)
	for _, a := range arr {
		endCount[a] = max(endCount[a], 1, endCount[a-difference]+1)
		ans = max(ans, endCount[a])
	}
	return ans
}

func max(nums ...int) int {
	a := nums[0]
	for _, i := range nums[1:] {
		if i > a {
			a = i
		}
	}
	return a
}
