package main

import "fmt"

func diagonal() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < len(matrix); i++ {
		// fmt.Print(matrix[i][i])
		fmt.Print(matrix[i][len(matrix)-1-i])
		// for j := 0; j < len(matrix); j++ {
		// 	// fmt.Print(matrix[i][j])
		// 	if i == j {
		// 		fmt.Print(matrix[i][j])
		// 	}

		// 	if i+j == len(matrix)-1 {
		// 		fmt.Print(matrix[i][j])
		// 	}
		// }
		fmt.Println()
	}
}

// Transpose the matrix
// Reverse each row
func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[j][i]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}

}

func spiralOrder(matrix [][]int) []int {
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	res := []int{}

	for left <= right && top <= bottom {
		// Print top row traverse left to right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// Print right col traverse top to bottom
		for i := right; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// Print bottom row traverse right to left
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}
		// Print left col traverse bottom to top
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}

func isValidSudokuV2(board [][]byte) bool {
	n := len(board)
	var rowUsed [9][9]bool
	var colUsed [9][9]bool
	var boxUsed [9][9]bool

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			ch := board[r][c]
			if ch == '.' {
				continue
			}

			digit := int(ch - '1') //Converts the character to an integer index (0–8 range)

			boxIndex := (r/3)*3 + (c / 3)

			if rowUsed[r][digit] || colUsed[c][digit] || boxUsed[boxIndex][digit] {
				return false
			}

			rowUsed[r][digit] = true
			colUsed[c][digit] = true
			boxUsed[boxIndex][digit] = true
		}
	}

	return true
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	firstZeroRow := false
	firstZeroCol := false

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstZeroRow = true
			break
		}
	}

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstZeroCol = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstZeroRow {
		for j := 1; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if firstZeroCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}

func main() {
	// diagonal()

	// Problem : Rotate Image (leecode 48)
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// rotate(matrix)
	// fmt.Println("Rotated Matrix:")
	// for _, row := range matrix {
	// 	fmt.Println(row)
	// }

	// Problem. Spiral Matrix (leecode 54)
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix1)) // [1 2 3 6 9 8 7 4 5]

	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix2)) // [1 2 3 4 8 12 11 10 9 5 6 7]

	// Problem: Valid Sudoku (leecode 36)
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudokuV2(board)) // true

	// Problem: Set Matrix Zeroes (leecode 73)

	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(matrix)
	fmt.Println(matrix)
}
