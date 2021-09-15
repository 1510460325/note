package main

import "fmt"

func main() {
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool, len(wordDict)+1)
	for _, w := range wordDict {
		set[w] = true
	}
	set[""] = true
	dp := make([]bool, len(s)+1)
	// dp[i] = dp[j] && set[s[j:i]]
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[j] && set[s[j:i]]
			if dp[i] {
				break
			}
		}
	}
	return dp[len(s)]
}
