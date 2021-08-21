package main

import "fmt"

func main() {
	fmt.Println(tribonacci(25))
}


func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	p1, p2, p3 := 0, 1, 1
	for i := 3; i <= n; i++ {
		p1, p2, p3 = p2, p3, p1 + p2 + p3
	}
	return p3
}