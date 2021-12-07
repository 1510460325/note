package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	grid := [][]int{
		{1,1},
		{1,2},
	}
	fmt.Println(colorBorder(grid, 0, 0, 3))
}

type loc struct {
	x, y int
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	set := map[string]bool{}
	queue := []loc{{row, col}}
	oldColor := grid[row][col]
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		key := fmt.Sprintf("%v:%v", now.x, now.y)
		if set[key] {
			continue
		}
		set[key] = true
		if now.x > 0 && grid[now.x-1][now.y] == oldColor {
			queue = append(queue, loc{now.x - 1, now.y})
		}
		if now.y > 0 && grid[now.x][now.y-1] == oldColor {
			queue = append(queue, loc{now.x, now.y - 1})
		}
		if now.x < len(grid)-1 && grid[now.x+1][now.y] == oldColor {
			queue = append(queue, loc{now.x + 1, now.y})
		}
		if now.y < len(grid[0])-1 && grid[now.x][now.y+1] == oldColor {
			queue = append(queue, loc{now.x, now.y + 1})
		}
	}
	isBorder := func(x, y int) bool {

		if x == 0 || x == len(grid) || y == 0 || y == len(grid[0]) {
			return true
		}
		return !set[fmt.Sprintf("%v:%v", x+1, y)] || !set[fmt.Sprintf("%v:%v", x, y+1)] ||
			!set[fmt.Sprintf("%v:%v", x-1, y)] || !set[fmt.Sprintf("%v:%v", x, y-1)]
	}
	for key := range set {
		param := strings.Split(key, ":")
		x, _ := strconv.Atoi(param[0])
		y, _ := strconv.Atoi(param[1])
		if isBorder(x, y) {
			grid[x][y] = color
		}
	}
	return grid
}
