package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/"))
}

func simplifyPath(path string) (ans string) {
	for strings.Contains(path, "//") {
		path = strings.ReplaceAll(path, "//", "/")
	}
	if len(path) == 0 {
		return "/"
	}
	if path[0] == '/' {
		path = path[1:]
	}
	if len(path) == 0 {
		return "/"
	}
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	router := strings.Split(path, "/")
	fmt.Println(router)
	stack := make([]string, 0)
	for _, item := range router {
		if item == "." {
			continue
		}
		if item == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, item)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	for _, item := range stack {
		ans += "/" + item
	}
	return ans
}
