package main

import "fmt"

func main() {
	fmt.Println(qPow(2, 10))
}

func qPow(a, n int) int {
	ans := 1
	for n > 0 {
		if n&1 == 1 {
			ans *= a
		}
		a *= a
		n >>= 1
	}
	return ans
}
