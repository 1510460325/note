package dijkstra

func networkDelayTime(times [][]int, n int, k int) (ans int) {
	graph := make([][]int, n)
	inf := 1<<31 - 1
	for s := 0; s < n; s++ {
		graph[s] = make([]int, n)
		for e := 0; e < n; e++ {
			if s == e {
				graph[s][e] = 0
			} else {
				graph[s][e] = inf
			}
		}
	}
	for _, t := range times {
		s, e := t[0]-1, t[1]-1
		graph[s][e] = t[2]
	}
	dist := make([]int, n)
	for s := 0; s < n; s++ {
		dist[s] = inf
	}
	dist[k-1] = 0
	done := make([]bool, n)
	for t := 0; t < n; t++ {
		now := -1
		for i, v := range done {
			if !v && (now == -1 || dist[i] < dist[now]) {
				now = i
			}
		}
		done[now] = true
		for e, l := range graph[now] {
			dist[e] = min(dist[e], dist[now] + l)
		}
	}
	for _, d := range dist {
		ans = max(ans, d)
	}
	if ans == inf {
		return -1
	}
	return ans
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
