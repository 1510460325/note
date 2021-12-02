package main

func main() {

}

func setZeroes(matrix [][]int) {
	zeroRow := make([]bool, len(matrix))
	zeroCol := make([]bool, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				zeroRow[i], zeroCol[j] = true, true
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if zeroRow[i] || zeroCol[j] {
				matrix[i][j] = 0
			}
		}
	}
}
