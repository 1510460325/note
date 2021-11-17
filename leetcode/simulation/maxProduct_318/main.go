package main

func main() {

}

func maxProduct(words []string) (ans int) {
	wInt := make([]int, len(words))
	for i := range wInt {
		for j := range words[i] {
			wInt[i] |= 1 << (words[i][j] - 'a')
		}
	}
	for i, iv := range wInt {
		for j, jv := range wInt {
			if iv&jv == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
