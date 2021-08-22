package main

import "fmt"

func main() {
	//fmt.Println(myAtoi("42"))
	//fmt.Println(myAtoi("   -42"))
	//fmt.Println(myAtoi("4193 with words"))
	//fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-88827   5655  U"))

}

func myAtoi(s string) (ans int) {
	start, negative := false, 1
	const maxValue = 1 << 31
	for i := range s {
		if s[i] == ' ' {
			if start {
				break
			}
		} else if s[i] == '+' || s[i] == '-' {
			if start {
				break
			}
			if s[i] == '-' {
				negative = -1
			}
			start = true
		} else if s[i] <= '9' && s[i] >= '0' {
			start = true
			now := s[i] - '0'
			if ans < maxValue/10 {
				ans = ans*10 + int(now)
				continue
			} else if ans == maxValue/10 {
				if negative == 1 && now >= (maxValue-1)%10 {
					ans = maxValue - 1
					break
				} else if negative == -1 && now >= maxValue%10 {
					ans = maxValue
					break
				} else {
					ans = ans*10 + int(now)
				}
			} else {
				if negative == 1 {
					ans = maxValue - 1
				} else {
					ans = maxValue
				}
				break
			}
		} else {
			break
		}
	}
	return ans * negative
}
