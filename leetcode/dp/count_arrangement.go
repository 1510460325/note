package main

import "fmt"

func main() {
	fmt.Println(shortestPathLength([][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}))
}

func shortestPathLength(graph [][]int) (ans int) {
	inf := 1<<31 - 2
	ans = inf
	n := len(graph)
	dp := make([][]int, 1<<n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for i := 0; i < n; i++ {
		dp[1<<i][i] = 0
	}
	// 最短距离
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = inf
		}
	}
	for v, nodes := range graph {
		for _, u := range nodes {
			d[v][u] = 1
		}
	}
	for k := range d {
		for i := range d {
			for j := range d {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	for mask := 1; mask < len(dp); mask++ {
		if mask&(mask-1) == 0 {
			continue
		}
		for u := 0; u < n; u++ {
			if mask>>u&1 > 0 {
				for v := 0; v < n; v++ {
					if v != u && mask>>v&1 > 0 {
						dp[mask][u] = min(dp[mask][u], dp[mask^(1<<u)][v]+d[v][u])
					}
				}
			}
		}
	}
	for _, i := range dp[1<<n-1] {
		ans = min(i, ans)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
