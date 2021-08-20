package main

func main() {

}

func minEatingSpeed(piles []int, h int) (ans int) {
	l, r := 1, -1
	for _, i := range piles {
		if r < i {
			r = i
		}
	}
	ans = r
	if canEatSuccess(piles, l, h) {
		return l
	}
	for l <= r {
		middle := (l + r) / 2
		if canEatSuccess(piles, middle, h) {
			ans = min(ans, middle)
			r = middle - 1
			if !canEatSuccess(piles, r, h) {
				return ans
			}
		} else {
			l = middle + 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func canEatSuccess(piles []int, speed, h int) bool {
	cost := 0
	for _, i := range piles {
		cost += i / speed
		if i%speed != 0 {
			cost++
		}
	}
	return cost <= h
}
