package main

func main() {

}

func arrangeCoins(n int) (ans int) {
	for i := 0; ;i++ {
		if i * (i+1) > 2 * n {
			return ans
		}
		ans = i
	}
	return ans
}