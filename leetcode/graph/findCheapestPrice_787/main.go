package main

func main() {

}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf int = 1<<31 + 1
	distance := make([][]int, n)
	for i := range distance {
		distance[i] = make([]int, n)
		for j := range distance[i] {
			distance[i][j] = inf
		}
		distance[i][i] = 0
	}
	for _, item := range flights {
		s, e, cost := item[0], item[1], item[2]
		distance[s][e] = cost
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
		dp[i][0] = distance[src][i]
	}
	for j := 1; j <= k; j++ {
		for i := 0; i < n; i++ {
			for u := 0; u < n; u++ {
				if u == i {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[u][j-1]+distance[u][i])
			}
		}
	}
	ans := inf
	for _, i := range dp[dst] {
		ans = min(ans, i)
	}
	if ans == inf {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
