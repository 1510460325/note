package main

import "sort"

func main() {

}

func findRadius(houses []int, heaters []int) (ret int) {
	sort.Ints(heaters)
	for _, i := range houses {
		ret = max(findMin(i, heaters), ret)
	}
	return ret
}

func findMin(now int, h []int) (ret int) {
	for idx := range h {
		if h[idx] >= now {
			if idx > 0 {
				return min(now-h[idx-1], h[idx]-now)
			}
			return h[idx] - now
		}
	}
	return now - h[len(h)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
