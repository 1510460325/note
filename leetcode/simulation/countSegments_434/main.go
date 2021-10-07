package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSegments("a"))
}

func countSegments(s string) (ans int) {
	s += " "
	valid := false
	for _, i := range []rune(s) {
		if i == ' ' {
			if valid {
				valid = false
				ans++
			}
			continue
		}
		valid = true
	}
	return ans
}
