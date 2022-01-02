package main

import "fmt"

func main() {
	fmt.Println(lastRemaining(2))
}

/*
1,2,3,4,5,6,7,8,9,10,11,12,13,14 1,14,1,true
2,4,6,8,10,12,14 2,14,2,false
4,8,12 4,12,4,true
8,8,8,false
*/
func lastRemaining(n int) int {
	s, e, c := 1, n, 1
	for s != e {
		s += c
		if (e-s)%(2*c) != 0 {
			e -= c
		}
		c *= 2
		if e == s {
			break
		}
		e -= c
		if (e-s)%(2*c) != 0 {
			s += c
		}
		c *= 2
	}
	return s
}
