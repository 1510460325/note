package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	for i := range board {
		if !validate(board[i]) {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		cur := make([]byte, 0, 9)
		for i := 0; i < 9; i++ {
			cur = append(cur, board[i][j])
		}
		if !validate(cur) {
			return false
		}
	}
	cur := make([]byte, 0, 9)
	cur = append(cur, board[0][:3]...)
	cur = append(cur, board[1][:3]...)
	cur = append(cur, board[2][:3]...)
	if !validate(cur) {
		return false
	}
	cur = make([]byte, 0, 9)
	cur = append(cur, board[0][3:6]...)
	cur = append(cur, board[1][3:6]...)
	cur = append(cur, board[2][3:6]...)
	if !validate(cur) {
		return false
	}
	cur = make([]byte, 0, 9)
	cur = append(cur, board[0][6:]...)
	cur = append(cur, board[1][6:]...)
	cur = append(cur, board[2][6:]...)
	if !validate(cur) {
		return false
	}
	cur = make([]byte, 0, 9)
	cur = append(cur, board[3][:3]...)
	cur = append(cur, board[4][:3]...)
	cur = append(cur, board[5][:3]...)
	if !validate(cur) {
		return false
	}
	cur = make([]byte, 0, 9)
	cur = append(cur, board[3][3:6]...)
	cur = append(cur, board[4][3:6]...)
	cur = append(cur, board[5][3:6]...)
	if !validate(cur) {
		return false
	}
	cur = make([]byte, 0, 9)
	cur = append(cur, board[3][6:]...)
	cur = append(cur, board[4][6:]...)
	cur = append(cur, board[5][6:]...)
	if !validate(cur) {
		return false
	}
	cur = make([]byte, 0, 9)
	cur = append(cur, board[6][:3]...)
	cur = append(cur, board[7][:3]...)
	cur = append(cur, board[8][:3]...)
	if !validate(cur) {
		return false
	}
	cur = make([]byte, 0, 9)
	cur = append(cur, board[6][3:6]...)
	cur = append(cur, board[7][3:6]...)
	cur = append(cur, board[8][3:6]...)
	if !validate(cur) {
		return false
	}
	cur = make([]byte, 0, 9)
	cur = append(cur, board[6][6:]...)
	cur = append(cur, board[7][6:]...)
	cur = append(cur, board[8][6:]...)
	if !validate(cur) {
		return false
	}
	return true
}

func validate(nums []byte) bool {
	used := make(map[byte]bool)
	for _, b := range nums {
		if b == '.' {
			continue
		}
		if used[b] {
			return false
		}
		used[b] = true
	}
	return true
}
