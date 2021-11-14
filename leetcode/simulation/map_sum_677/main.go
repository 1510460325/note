package main

import "strings"

func main() {

}

type MapSum map[string]int

func Constructor() MapSum {
	return MapSum{}
}

func (s *MapSum) Insert(key string, val int) {
	(*s)[key] = val
}

func (s *MapSum) Sum(prefix string) (ans int) {
	for k, v := range *s {
		if strings.HasPrefix(k, prefix) {
			ans += v
		}
	}
	return ans
}
