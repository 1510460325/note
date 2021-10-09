package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(0))
}

func isPalindrome(x int) bool {
	y := 0
	a := x
	for a > 0 {
		y = y*10 + a%10
		a /= 10
	}
	return y == x
}
