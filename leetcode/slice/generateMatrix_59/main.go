package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(1))
}

type loc struct {
	x, y int
}

func generateMatrix(n int) (graph [][]int) {
	graph = make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}
	a, b := loc{0, 0}, loc{n - 1, n - 1}
	idx := 0
	nextId := func() int {
		idx++
		return idx
	}
	for a.x <= b.x {
		for j := a.y; j <= b.y; j++ {
			graph[a.x][j] = nextId()
		}
		for i := a.x+1; i <= b.x; i++ {
			graph[i][b.y] = nextId()
		}
		for j := b.y - 1; j >= a.y; j-- {
			graph[b.x][j] = nextId()
		}
		for i := b.x - 1; i > a.x; i-- {
			graph[i][a.y] = nextId()
		}
		a.x++
		a.y++
		b.x--
		b.y--
	}
	return graph
}
