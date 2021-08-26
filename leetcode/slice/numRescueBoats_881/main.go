package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{2,4}, 5))
}

func numRescueBoats(people []int, limit int) (ans int) {
	sort.Ints(people)
	l, r := 0, len(people)-1
	for l < r {
		ans++
		if people[r]+people[l] <= limit {
			l++
		}
		r--
	}
	if l == r {
		ans++
	}
	return ans
}
