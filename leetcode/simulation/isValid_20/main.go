package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]]{}"))
}

func isValid(s string) bool {
	stack := make([]int32, 0)
	for _, i := range s {
		switch i {
		case '[', '(', '{':
			stack = append(stack, i)
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
