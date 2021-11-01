package main

import (
	"fmt"
)

func main() {
	fmt.Println(distributeCandies([]int{}))
}

func distributeCandies(candyType []int) int {
	typeMap := make(map[int]bool)
	for _, t := range candyType {
		typeMap[t] = true
	}
	if len(typeMap) > len(candyType)/2 {
		return len(candyType) / 2
	}
	return len(typeMap)
}
