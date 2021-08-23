package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(10))
}

func intToRoman(num int) (ans string) {
	// 个位
	if num == 0 {
		return ans
	}
	ans = getRoman("I", "V", "X", num%10)
	num /= 10
	// 十位
	if num == 0 {
		return ans
	}
	ans = getRoman("X", "L", "C", num%10) + ans
	num /= 10
	// 百位
	if num == 0 {
		return ans
	}
	ans = getRoman("C", "D", "M", num%10) + ans
	num /= 10
	// 千位
	if num == 0 {
		return ans
	}
	a := num % 10
	for i := 1; i <= a; i++ {
		ans = "M" + ans
	}
	return ans
}

func getRoman(one, five, ten string, a int) (ans string) {
	if 1 <= a && a <= 3 {
		for i := 1; i <= a; i++ {
			ans += one
		}
	} else if a == 4 {
		ans = one + five
	} else if a == 5 {
		ans = five
	} else if 5 < a && a < 9 {
		ans = five
		for i := 6; i <= a; i++ {
			ans += one
		}
	} else if a == 9 {
		ans = one + ten
	}
	return ans
}
