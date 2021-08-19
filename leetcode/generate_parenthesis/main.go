package main

import "fmt"

func main() {
	fmt.Println(len(generateParenthesis(4)))
	ans := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}
	fmt.Println(len(ans))
}

func generateParenthesis(n int) (ans []string) {
	dp := make([][]string, n+1)
	dp[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		m := make(map[string]bool)
		for j := 1; j < i; j++ {
			a, b := genBorder(j)
			for _, k := range dp[i-j] {
				for _, s := range dp[j] {
					m[s+k] = true
					m[a+k+b] = true
				}
			}
		}
		for k := range m {
			dp[i] = append(dp[i], k)
		}
	}
	return dp[n]
}

func genBorder(a int) (b string, c string) {
	for i := 0; i < a; i++ {
		b += "("
		c += ")"
	}
	return b, c
}
