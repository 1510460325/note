package main

import "fmt"

func main() {
	fmt.Println(checkRecordA(10101))
}

func checkRecord(s string) bool {
	absentTime, sequencesLateTime := 0, 0
	for i := range s {
		if s[i] == 'A' {
			absentTime++
			if absentTime == 2 {
				return false
			}
		} else if s[i] == 'L' {
			sequencesLateTime++
			if sequencesLateTime == 3 {
				return false
			}
		} else {
			sequencesLateTime = 0
		}
	}
	return true
}

// 552. 学生出勤记录 II
func checkRecordA(n int) (ans int) {
	if n == 0 {
		return 0
	}
	const mod int = 1e9 + 7
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, 2)
		dp[i][0] = make([]int, 3)
		dp[i][1] = make([]int, 3)
	}
	dp[0][1][0] = 1
	dp[0][0][1] = 1
	dp[0][0][0] = 1
	for i := 1; i < n; i++ {
		// Absent
		for k := 0; k <= 2; k++ {
			dp[i][1][0] = (dp[i][1][0] + dp[i-1][0][k]) % mod
		}
		// Present
		for j := 0; j < 2; j++ {
			for k := 0; k <= 2; k++ {
				dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][k]) % mod
			}
		}
		// Late
		for j := 0; j < 2; j++ {
			for k := 1; k <= 2; k++ {
				dp[i][j][k] = (dp[i][j][k] + dp[i-1][j][k-1]) % mod
			}
		}
	}
	for j := 0; j < 2; j++ {
		for k := 0; k < 3; k++ {
			ans = (ans + dp[n-1][j][k]) % mod
		}
	}
	return ans
}
