package main

import "fmt"

func main() {
	fmt.Println(multiply("0", "52"), 123456789*987654321)
	fmt.Println(123456789 * 9)
}

func multiply(num1 string, num2 string) (ans string) {
	if num2 == "0" || num1 == "0" {
		return "0"
	}
	list := make([][]int, 0, len(num2))
	prefix := make([]int, 0)
	for i := len(num2) - 1; i >= 0; i-- {
		c, b := 0, int(num2[i]-'0')
		item := make([]int, len(prefix))
		copy(item, prefix)
		prefix = append(prefix, 0)
		if b == 0 {
			list = append(list, []int{})
			continue
		}
		for j := len(num1) - 1; j >= 0; j-- {
			a := int(num1[j] - '0')
			m := a*b + c
			item = append(item, m%10)
			c = m / 10
		}
		if c != 0 {
			item = append(item, c)
		}
		list = append(list, item)
	}
	c := 0
	for i := 0; ; i++ {
		m := c
		end := true
		for _, item := range list {
			if len(item) <= i {
				continue
			} else {
				end = false
				m += item[i]
			}
		}
		if end && c == 0 {
			break
		}
		fmt.Println(i, m)
		ans = string(rune('0'+m%10)) + ans
		if end {
			break
		}
		c = m / 10
	}
	if ans == "" {
		return "0"
	}
	return ans
}
