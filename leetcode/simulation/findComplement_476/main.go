package main

import "fmt"

func main() {
	fmt.Println(findComplement(1))
}
func findComplement(num int) (ans int) {
	base := 1
	for num > 0 {
		if num&1 == 0 {
			ans += base
		}
		num >>= 1
		base <<= 1
	}
	return ans
}
