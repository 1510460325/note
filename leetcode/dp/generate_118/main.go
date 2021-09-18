package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = 1
	for i := 1; i < numRows; i++ {
		dp[i][0] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		}
		dp[i][i] = 1
	}
	return dp
}

func getRow(rowIndex int) []int {
	return generate(rowIndex + 1)[rowIndex]
}
