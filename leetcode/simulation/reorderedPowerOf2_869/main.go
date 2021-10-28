package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(reorderedPowerOf2(46))
}

func reorderedPowerOf2(n int) bool {
	in := []byte(strconv.Itoa(n))
	sort.Slice(in, func(i, j int) bool {
		return in[i] > in[j]
	})
	maxIn, _ := strconv.Atoi(string(in))
	a := 1
	for {
		if a > maxIn {
			break
		}
		base := []byte(strconv.Itoa(a))
		sort.Slice(base, func(i, j int) bool {
			return base[i] > base[j]
		})
		if string(base) == string(in) {
			return true
		}
		a *= 2
	}
	return false
}
