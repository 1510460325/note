package main

import "fmt"

func main() {
	fmt.Println(validTicTacToe([]string{"OXX", "XOX", "OXO"}))
}

func validTicTacToe(board []string) bool {
	success := func(cur uint8) bool {
		return board[0][0] == cur && board[0][1] == cur && board[0][2] == cur ||
			board[1][0] == cur && board[1][1] == cur && board[1][2] == cur ||
			board[2][0] == cur && board[2][1] == cur && board[2][2] == cur ||
			board[0][0] == cur && board[1][0] == cur && board[2][0] == cur ||
			board[0][1] == cur && board[1][1] == cur && board[2][1] == cur ||
			board[0][2] == cur && board[1][2] == cur && board[2][2] == cur ||
			board[0][0] == cur && board[1][1] == cur && board[2][2] == cur ||
			board[0][2] == cur && board[1][1] == cur && board[2][0] == cur
	}
	countValidate := func() bool {
		a, b := 0, 0
		for _, s := range board {
			for _, i := range s {
				if i == 'X' {
					a++
				} else if i == 'O' {
					b++
				}
			}
		}
		return a == b && !success('X') || a == b+1 && !success('O')
	}
	return countValidate()
}
