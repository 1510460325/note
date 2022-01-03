package main

import "fmt"

func main() {
	fmt.Println(partition("a"))
}

func partition(s string) (rst [][]string) {
	var dfs func(pre []string, cur string)
	dfs = func(pre []string, cur string) {
		if len(cur) == 0 {
			tmp := make([]string, len(pre))
			copy(tmp, pre)
			rst = append(rst, tmp)
			return
		}
		for i := 1; i <= len(cur); i++ {
			if isValid(cur[:i]) {
				dfs(append(pre, cur[:i]), cur[i:])
			}
		}
	}
	dfs(nil, s)
	return rst
}

func isValid(s string) bool {
	for i := 0; i < len(s) && i <= len(s)-i-1; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
