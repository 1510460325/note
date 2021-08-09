package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(nthUglyNumber(10))
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	sort.Ints(primes)
	points := make([]int, len(primes))
	for i := range primes {
		points[i] = 0
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		minIndex := 0
		for j := range primes {
			for primes[j] * dp[points[j]] <= dp[i-1] {
				points[j]++
			}
			now := primes[j] * dp[points[j]]
			if now < primes[minIndex] * dp[points[minIndex]] {
				minIndex = j
			}
		}
		dp[i] = primes[minIndex] * dp[points[minIndex]]
		points[minIndex]++
	}
	return dp[n-1]
}
