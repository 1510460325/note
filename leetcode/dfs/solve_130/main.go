package main

func main() {

}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	found := make([][]int, m)
	for i := range found {
		found[i] = make([]int, n)
	}

	var escape func(i, j int, step int) bool

	escape = func(i, j int, step int) (rst bool) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if found[i][j] > 0 || board[i][j] == 'X' {
			return false
		}
		found[i][j] = step
		rst = i == 0 || i == m-1 || j == 0 || j == n-1
		rst = escape(i-1, j, step) || rst
		rst = escape(i+1, j, step) || rst
		rst = escape(i, j-1, step) || rst
		rst = escape(i, j+1, step) || rst
		return rst
	}

	step := 0
	kill := make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			step++
			if !escape(i, j, step) {
				kill[step] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if kill[found[i][j]] {
				board[i][j] = 'X'
			}
		}
	}
}
