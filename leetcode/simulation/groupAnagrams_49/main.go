package main

import "sort"

func main() {

}

func groupAnagrams(strList []string) (ans [][]string) {
	m := make(map[string][]string)
	for _, s := range strList {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedStr := string(runes)
		m[sortedStr] = append(m[sortedStr], s)
	}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}