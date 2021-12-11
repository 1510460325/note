package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(restoreIpAddresses("010010"))
}

func strToInt(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func isValid(s string) bool {
	now := strToInt(s)
	return now >= 0 && now <= 255 && (len(s) == 1 || s[0] != '0')
}

func restoreIpAddresses(s string) (ret []string) {
	var dfs func(str string, preIdx, rest int)
	dfs = func(str string, preIdx, rest int) {
		// end
		if rest == 0 {
			if isValid(str[preIdx+1:]) {
				ret = append(ret, str)
			}
			return
		}
		// insert '.'
		for i := preIdx + 2; i < len(str); i++ {
			if isValid(str[preIdx+1:i]) {
				dfs(str[:i] +"."+str[i:], i, rest-1)
			} else {
				return
			}
		}
	}
	dfs(s, -1, 3)
	return ret
}
