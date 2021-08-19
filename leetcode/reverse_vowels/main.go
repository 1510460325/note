package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	index := make([]int, 0)
	for i := range s {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U'{
			index = append(index, i)
		}
	}
	ans := []uint8(s)
	if len(index) < 2 {
		return s
	}
	i, j := 0, len(index)-1
	for i < j {
		ans[index[i]], ans[index[j]] = ans[index[j]], ans[index[i]]
		i++
		j--
	}
	return string(ans)
}
