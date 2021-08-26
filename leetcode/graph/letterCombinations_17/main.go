package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) (ans []string) {
	if len(digits) == 0 {
		return ans
	}
	graph := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
	var dfs func(int)
	var path []byte
	dfs = func(step int) {
		if step == len(digits) {
			ans = append(ans, string(path))
			return
		}
		for _, now := range graph[digits[step]-'2'] {
			path = append(path, now)
			dfs(step + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
