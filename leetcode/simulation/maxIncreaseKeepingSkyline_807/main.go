package main

func main() {

}

func maxIncreaseKeepingSkyline(grid [][]int) (ret int) {
	rowMax := make([]int, len(grid))
	for i := range grid {
		for _, cur := range grid[i] {
			rowMax[i] = max(rowMax[i], cur)
		}
	}
	colMax := make([]int, len(grid[0]))
	for j := range grid[0] {
		for i := range grid {
			colMax[j] = max(colMax[j], grid[i][j])
		}
	}
	for i := range grid {
		for j := range grid[i] {
			ret += min(rowMax[i], colMax[j]) - grid[i][j]
		}
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}