package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	c := Constructor([]int{1, 2, 3})
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
	fmt.Println(c.Shuffle())
}

type Solution struct {
	data     []int
	now      []int
	maxCount int
	count    int
}

func Constructor(nums []int) Solution {
	a := make([]int, len(nums))
	copy(a, nums)
	return Solution{
		now:      nums,
		data:     a,
		maxCount: int(math.Pow(2, float64(len(nums)))),
		count:    0,
	}
}

func (s *Solution) Reset() []int {
	return s.data
}

func (s *Solution) Shuffle() []int {
	s.count = (s.count + 1) % s.maxCount
	s.next()
	if s.count == 0 {
		return s.Shuffle()
	}
	return s.now
}

func (s *Solution) next() {
	num := s.now
	for i := len(num) - 2; i >= 0; i-- {
		if num[i] < num[i+1] {
			for j := len(num) - 1; j > i; j-- {
				if num[j] > num[i] {
					num[j], num[i] = num[i], num[j]
					sort.Ints(num[i+1:])
					return
				}
			}
		}
	}
	sort.Ints(num)
}
