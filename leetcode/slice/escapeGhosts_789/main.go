package main

func main() {

}

func escapeGhosts(ghosts [][]int, target []int) bool {
	ghostMin := 1<<63 - 1
	for _, ghost := range ghosts {
		ghostMin = min(manhattanDistance(ghost, target), ghostMin)
	}
	return ghostMin > manhattanDistance(target, []int{0, 0})
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func manhattanDistance(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}
