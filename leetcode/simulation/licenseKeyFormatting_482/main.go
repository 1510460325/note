package main

import "fmt"

func main() {
	fmt.Println(licenseKeyFormatting("a-a-a-a--", 2))
}

func licenseKeyFormatting(s string, k int) (ans string) {
	count := 0
	gap := byte('z' - 'Z')
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		if count == k {
			ans = "-" + ans
			count = 0
		}
		if s[i] <= 'z' && s[i] >= 'a' {
			ans = string(s[i]-gap) + ans
		} else {
			ans = string(s[i]) + ans
		}
		count++
	}
	return ans
}
