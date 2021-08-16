package main

func main() {
}

func maxArea(height []int) (ans int) {
	l, r := 0, len(height)-1
	for l < r {
		now := min(height[l], height[r]) * (r - l)
		ans = max(now, ans)
		if height[l] < height[r] {
			l++
		} else {
			r--
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
