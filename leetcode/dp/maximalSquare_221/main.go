package main

import "fmt"

func main() {
	fmt.Println(maximalSquareV2([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	matSum := make([][]int, m)
	for i := range matSum {
		matSum[i] = make([]int, n)
	}
	for i := range matSum {
		temSum := 0
		for j := range matSum[i] {
			temSum += int(matrix[i][j] - '0')
			if i == 0 {
				matSum[i][j] = temSum
			} else {
				matSum[i][j] = matSum[i-1][j] + temSum
			}
		}
	}
	getValue := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return matSum[i][j]
	}
	squareArea := func(i0, j0, i1, j1 int) int {
		value := getValue(i1, j1) + getValue(i0-1, j0-1) - getValue(i0-1, j1) - getValue(i1, j0-1)
		if value == (i1-i0+1)*(i1-i0+1) {
			return value
		}
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; i-k >= 0 && j-k >= 0; k++ {
				value := squareArea(i-k, j-k, i, j)
				if value > 0 {
					dp[i][j] = value
				} else {
					break
				}
			}
			if i > 0 && dp[i-1][j] > dp[i][j] {
				dp[i][j] = dp[i-1][j]
			}
			if j > 0 && dp[i][j-1] > dp[i][j] {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func maximalSquareV2(matrix [][]byte) (ans int) {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			a, b, c := 0, 0, 0
			if i > 0 && j > 0 {
				a = dp[i-1][j-1]
			}
			if i > 0 {
				b = dp[i-1][j]
			}
			if j > 0 {
				c = dp[i][j-1]
			}
			dp[i][j] = min(a, b, c) + 1
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans * ans
}

func min(nums ...int) int {
	tem := nums[0]
	for _, i := range nums {
		if i < tem {
			tem = i
		}
	}
	return tem
}
