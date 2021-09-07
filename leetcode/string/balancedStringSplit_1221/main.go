package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("LLLLRRRR"))
}

func balancedStringSplit(s string) (ans int) {
	lCount, rCount := 0, 0
	for i := range s {
		if s[i] == 'R' {
			rCount++
		} else {
			lCount++
		}
		if lCount == rCount {
			ans++
			lCount, rCount = 0, 0
		}
	}
	return ans
}
