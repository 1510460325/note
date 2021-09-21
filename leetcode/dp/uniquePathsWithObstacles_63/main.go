package main

func main() {

}

func uniquePathsWithObstacles(graph [][]int) int {
	m, n := len(graph), len(graph[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = (graph[0][0] + 1) % 2
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 1 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
