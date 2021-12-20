package main

import "sort"

func main() {

}

func findRadius(houses []int, heaters []int) (ret int) {
	sort.Ints(heaters)
	sort.Ints(houses)
	m, start := 0, 0
	for _, i := range houses {
		m, start = findMin(i, heaters, start)
		ret = max(m, ret)
	}
	return ret
}

func findMin(now int, h []int, start int) (_, _ int) {
	for idx := start; idx < len(h); idx++ {
		if h[idx] >= now {
			if idx > 0 {
				return min(now-h[idx-1], h[idx]-now), idx
			}
			return h[idx] - now, idx
		}
	}
	return now - h[len(h)-1], len(h) - 1
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
