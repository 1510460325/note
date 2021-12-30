package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	numCount, group := make([][]int, 0, len(hand)), len(hand)/groupSize
	for _, item := range hand {
		if len(numCount) == 0 || numCount[len(numCount)-1][0] != item {
			numCount = append(numCount, []int{item, 1})
		} else {
			numCount[len(numCount)-1][1]++
		}
	}
	cut := func() bool {
		if len(numCount) < groupSize {
			return false
		}
		for i := 0; i < groupSize; i++ {
			if i == 0 || numCount[i][0] == numCount[i-1][0]+1 {
				numCount[i][1]--
			} else {
				return false
			}
		}
		for i := 0; i < groupSize && i < len(numCount); i++ {
			if numCount[i][1] == 0 {
				numCount = append(numCount[0:i], numCount[i+1:]...)
				i--
			}
		}
		return true
	}
	for i := 0; i < group; i++ {
		if !cut() {
			return false
		}
	}
	return true
}
