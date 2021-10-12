package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IV"))
}

func romanToInt(s string) (ans int) {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			ans += 1
			if i+1 < len(s) && s[i+1] == 'V' {
				ans += 3
				i++
			}
			if i+1 < len(s) && s[i+1] == 'X' {
				ans += 8
				i++
			}
		case 'X':
			ans += 10
			if i+1 < len(s) && s[i+1] == 'L' {
				ans += 30
				i++
			}
			if i+1 < len(s) && s[i+1] == 'C' {
				ans += 80
				i++
			}
		case 'C':
			ans += 100
			if i+1 < len(s) && s[i+1] == 'D' {
				ans += 300
				i++
			}
			if i+1 < len(s) && s[i+1] == 'M' {
				ans += 800
				i++
			}
		case 'V':
			ans += 5
		case 'L':
			ans += 50
		case 'D':
			ans += 500
		case 'M':
			ans += 1000
		}
	}
	return ans
}
