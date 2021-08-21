package main

import "fmt"

func main() {
}

func compress(chars []byte) (ans int) {
	if len(chars) == 0 {
		return ans
	}
	pre, count, index, n := chars[0], 1, 1, len(chars)
	for i := 1; i < n; i++ {
		if chars[i] == pre {
			count++
			continue
		}
		if count > 1 {
			for _, j := range []byte(getStr(count)) {
				chars[index] = j
				index++
			}
		}
		pre = chars[i]
		chars[index] = pre
		index++
		count = 1
	}
	if count > 1 {
		for _, j := range []byte(getStr(count)) {
			chars[index] = j
			index++
		}
	}
	return index
}

func getStr(n int) (s string) {
	for n > 0 {
		s = fmt.Sprintf("%v", n%10) + s
		n /= 10
	}
	return s
}
