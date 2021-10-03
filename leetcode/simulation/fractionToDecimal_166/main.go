package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(1, 333))
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	ans := make([]byte, 0)
	if (numerator < 0) != (denominator < 0) {
		ans = append(ans, '-')
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	// 整数部分
	ans = append(ans, strconv.Itoa(numerator/denominator)...)
	ans = append(ans, '.')
	// 小数部分
	lastIndex := make(map[int]int)
	last := (numerator % denominator) * 10
	decimal := make([]byte, 0)
	index, ok := -1, false
	for i := 0; ; i++ {
		index, ok = lastIndex[last]
		if ok {
			break
		}
		cur := last / denominator
		decimal = append(decimal, strconv.Itoa(cur)...)
		if last%denominator == 0 {
			break
		}
		lastIndex[last] = i
		last = last % denominator * 10
	}
	if last%denominator == 0 {
		ans = append(ans, decimal...)
		return string(ans)
	}
	// 循环
	ans = append(ans, decimal[:index]...)
	ans = append(ans, '(')
	ans = append(ans, decimal[index:]...)
	ans = append(ans, ')')
	return string(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
