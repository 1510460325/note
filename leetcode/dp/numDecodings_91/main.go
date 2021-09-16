package main

import "fmt"

func main() {
	fmt.Println(numDecodings("0"))
}

func numDecodings(s string) int {
	dp := make([]int, len(s))
	for i := range s {
		dp[i] = getNum(s[:i+1])
		if i >= 1 {
			dp[i] += dp[i-1] * getNum(s[i:i+1])
		}
		if i >= 2 {
			dp[i] += dp[i-2] * getNum(s[i-1:i+1])
		}
	}
	return dp[len(s)-1]
}

func getNum(s string) int {
	if canDecode(s) {
		return 1
	}
	return 0
}

func canDecode(s string) bool {
	switch len(s) {
	case 1:
		return s[0] != '0'
	case 2:
		return s[0] == '1' || (s[0] == '2' && s[1] <= '6')
	default:
		return false
	}
}
