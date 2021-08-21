package main

import "fmt"

func main() {
	fmt.Println(convert("AB", 1))
}

func reverseVowels(s string) string {
	index := make([]int, 0)
	for i := range s {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			index = append(index, i)
		}
	}
	ans := []uint8(s)
	if len(index) < 2 {
		return s
	}
	i, j := 0, len(index)-1
	for i < j {
		ans[index[i]], ans[index[j]] = ans[index[j]], ans[index[i]]
		i++
		j--
	}
	return string(ans)
}

func reverseStr(s string, k int) (ans string) {
	for i := 0; i < len(s); i += 2 * k {
		j := i + 2*k
		if j <= len(s) {
			ans += reverseK(s[i:i+k]) + s[i+k:j]
			continue
		}
		j = i + k
		if j <= len(s) {
			ans += reverseK(s[i:i+k]) + s[i+k:]
			continue
		}
		ans += reverseK(s[i:])
	}
	return ans
}

func reverseK(s string) (ans string) {
	for _, j := range s {
		ans = string(j) + ans
	}
	return ans
}

func convert(s string, numRows int) (ans string) {
	if numRows == 1 {
		return s
	}
	sl := make([]string, numRows)
	r, down := 0, false
	for _, i := range s {
		sl[r] += string(i)
		if r == 0 || r == numRows-1 {
			down = !down
		}
		if down {
			r++
		} else {
			r--
		}
	}
	for _, i := range sl {
		ans += i
	}
	return ans
}
