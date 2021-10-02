package main

import "fmt"

func main() {
	fmt.Println(toHex(14))
}

func toHex(num int) (hex string) {
	flag := true
	if num < 0 {
		num = -num
		flag = false
	}
	if num == 0 {
		return "0"
	}
	hexNum, idx := [8]int{}, 7
	numHex := [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	for num > 0 {
		a := num % 16
		hex = numHex[a] + hex
		num /= 16
		hexNum[idx], idx = a, idx-1
	}
	if flag {
		return hex
	}
	c := 1
	for i := 7; i >= 0; i-- {
		hexNum[i] = 15 - hexNum[i] + c
		c, hexNum[i] = hexNum[i]/16, hexNum[i]%16
	}
	hex = ""
	start := false
	for _, i := range hexNum {
		if i == 0 && !start {
			continue
		}
		start = true
		hex += numHex[i]
	}
	return hex
}
