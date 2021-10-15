package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	descNumber, next := "1", ""
	for i := 1; i < n; i++ {
		ch, count := '0', 0
		for _, c := range descNumber {
			if c == ch {
				count++
			} else {
				if count > 0 {
					next += fmt.Sprintf("%v%c", count, ch)
				}
				ch, count = c, 1
			}
		}
		if count > 0 {
			next += fmt.Sprintf("%v%c", count, ch)
		}
		descNumber, next = next, ""
	}
	return descNumber
}
