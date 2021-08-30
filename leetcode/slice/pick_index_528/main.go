package main

import (
	"math/rand"
)

func main() {

}

type Solution struct {
	sum      int
	maxIndex []int
}

func Constructor(w []int) Solution {
	maxIndex := make([]int, len(w))
	sum := 0
	for i, v := range w {
		sum += v
		maxIndex[i] = sum - 1
	}
	return Solution{
		sum:      sum,
		maxIndex: maxIndex,
	}
}

func (s *Solution) PickIndex() int {
	n := rand.Intn(s.sum)
	for v, max := range s.maxIndex {
		if max >= n {
			return v
		}
	}
	return -1
}
