package main

func main() {

}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] == target {
			return l
		} else if nums[r] == target {
			return r
		}
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		if target < nums[m] {
			if nums[l] < target || nums[m] < nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] > nums[l] || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
