package main

func main() {

}

func countBattleships(board [][]byte) (ret int) {
	var destroy func(x, y int)
	destroy = func(x, y int) {
		if x == len(board) || y == len(board[0]) || board[x][y] != 'X' {
			return
		}
		board[x][y] = '.'
		destroy(x+1, y)
		destroy(x, y+1)
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				ret++
				destroy(i, j)
			}
		}
	}
	return
}
