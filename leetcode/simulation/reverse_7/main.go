package main

import "fmt"

func main() {
	fmt.Println(reverse(-120))
}

func reverse(x int) (ans int) {
	c := 1
	if x < 0 {
		c = -1
		x = -x
	}
	const max = 1 << 31
	for x > 0 {
		if ans > max/10 {
			return 0
		}
		if ans > max/10 ||
			(c == -1 && (max+1)-ans*10 < x%10) ||
			(c == 1 && max-ans*10 < x%10) {
			return 0
		}
		ans = ans*10 + x%10
		x /= 10
	}
	return ans * c
}
