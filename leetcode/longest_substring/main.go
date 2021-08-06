package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("ab"))
}

func lengthOfLongestSubstring(s string) (ans int) {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	l, r := 0, 1
	used := make(map[uint8]bool)
	used[s[0]] = true
	for r != len(s) {
		if used[s[r]] {
			ans = max(r-l, ans)
			used[s[l]] = false
			l++
		} else {
			used[s[r]] = true
			r++
		}
	}
	ans = max(r-l, ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
