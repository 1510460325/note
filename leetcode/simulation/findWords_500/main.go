package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
func findWords(words []string) (ans []string) {
	set := make(map[uint8]int)
	str := "qwertyuiop"
	for i := range str {
		set[str[i]] = 1
	}
	str = "asdfghjkl"
	for i := range str {
		set[str[i]] = 2
	}
	str = "zxcvbnm"
	for i := range str {
		set[str[i]] = 3
	}
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		lower := strings.ToLower(w)
		for i := range lower {
			if set[lower[0]] != set[lower[i]] {
				goto fail
			}
		}
		ans = append(ans, w)
	fail:
	}
	return ans
}
