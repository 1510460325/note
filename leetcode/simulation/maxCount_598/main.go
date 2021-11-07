package main

func main() {

}

func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		a, b := op[0], op[1]
		if a < m {
			m = a
		}
		if b < n {
			n = b
		}
	}
	return m * n
}
