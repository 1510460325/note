package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCB"))
}

type loc struct {
	x, y int
}

func exist(board [][]byte, word string) bool {

	skip := func(path []loc, now loc) bool {
		for _, p := range path {
			if p.x == now.x && p.y == now.y {
				return true
			}
		}
		return false
	}

	var search func(path []loc, now loc, rest []byte) bool
	search = func(path []loc, now loc, rest []byte) bool {
		if len(rest) == 0 {
			return true
		}
		if rest[0] != board[now.x][now.y] {
			return false
		}
		locList := []loc{{now.x - 1, now.y}, {now.x + 1, now.y}, {now.x, now.y - 1}, {now.x, now.y + 1}}
		path = append(path, now)
		for _, next := range locList {
			if next.x < 0 || next.x >= len(board) || next.y < 0 || next.y >= len(board[0]) {
				continue
			}
			if skip(path, next) {
				continue
			}
			if search(path, next, rest[1:]) {
				return true
			}
		}
		path = path[:len(path)-1]
		return len(rest) == 1
	}
	byteList := []byte(word)
	fmt.Println(board, byteList)
	for i := range board {
		for j := range board[i] {
			if search(nil, loc{i, j}, byteList) {
				return true
			}
		}
	}
	return false
}
