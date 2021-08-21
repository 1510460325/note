package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 5, 7}))
}

func numberOfArithmeticSlices(nums []int) (ans int) {
	if len(nums) < 3 {
		return 0
	}
	s := 0
	doCount := func(a int) int {
		return (a - 2) * (a - 1) / 2
	}
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] != nums[i-1]-nums[i-2] {
			if i-s >= 3 {
				ans += doCount(i - s)
				fmt.Printf("once: %v, s: %v, i: %v\n", doCount(i-s), s, i)
			}
			s = i - 1
		} else if i == len(nums)-1 {
			if i-s+1 >= 3 {
				ans += doCount(i - s + 1)
				fmt.Printf("last: %v, s:%v, i: %v\n", doCount(i-s+1), s, i)
			}
		}
	}
	return ans
}

func numberOfArithmeticSlices1(nums []int) (ans int) {
	f := make([]map[int]int, len(nums))
	for i, x := range nums {
		f[i] = map[int]int{}
		for j, y := range nums[:i] {
			d := x - y
			cnt := f[j][d]
			ans += cnt
			f[i][d] += cnt + 1
		}
	}
	return
}
