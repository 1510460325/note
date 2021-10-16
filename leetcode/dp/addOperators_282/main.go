package main

func main() {

}

func addOperators(num string, target int) []string {
	n := len(num)
	dp := make([][]map[int][]string, n)
	for i := range dp {
		dp[i] = make([]map[int][]string, n)
		dp[i][i] = map[int][]string{int(num[i] - '0'): {num[i : i+1]}}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := i; k < j; k++ {
				for aVal, aList := range dp[i][k] {
					for bVal, bList := range dp[k+1][j] {
						for _, a := range aList{
							for _, b := range bList{
								dp[i][j][aVal+bVal] = append(dp[i][j][aVal+bVal], a + "+" + b)
								dp[i][j][aVal*bVal] = append(dp[i][j][aVal*bVal], a + "*" + b)
							}
						}

					}
				}
			}
		}
	}
	return nil
}
