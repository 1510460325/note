package main

func main() {

}

func maxProfit(prices []int) (ans int) {
	preMin := prices[0]
	for i := 1; i < len(prices); i++ {
		ans = max(ans, prices[i]-preMin)
		preMin = min(prices[i], preMin)
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
