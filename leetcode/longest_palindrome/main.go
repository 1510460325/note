package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	max := 0
	ans := []int{0, 0}
	for x := 2; x <= len(s); x++ {
		for l := 0; l < len(s); l++ {
			r := l + x - 1
			if r >= len(s) {
				break
			}
			if s[l] != s[r] {
				dp[l][r] = false
				continue
			}
			if x == 2 {
				dp[l][r] = true
			} else {
				dp[l][r] = dp[l+1][r-1]
			}
			if dp[l][r] && max < x {
				max = x
				ans = []int{l, r}
			}
		}
	}
	return s[ans[0] : ans[1]+1]
}
