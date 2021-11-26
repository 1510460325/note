package main

import "fmt"

func main() {
	fmt.Println(myPow(2, -2))
}

func myPow(x float64, n int) (ans float64) {
	if n == 0 {
		return 1
	}
	reverse := n < 0
	if reverse {
		n = -n
	}
	ans = 1
	for n > 0 {
		if n&1 > 0 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	if reverse {
		return 1 / ans
	}
	return ans
}
