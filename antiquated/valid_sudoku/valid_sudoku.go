package valid_sudoku

func IsValidSudokuTwo(board [][]byte) bool {
	black := make([][]bool, 9)
	for i := 0; i < len(black); i++ {
		black[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		rows, cols := make(map[byte]bool, 9), make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			blackIdx := i/3*3 + j/3
			if (board[i][j] != '.' && rows[board[i][j]]) || (board[j][i] != '.' && cols[board[j][i]]) ||
				(board[i][j] != '.' && black[blackIdx][board[i][j]-'1']) {
				return false
			}
			if board[i][j] != '.' {
				rows[board[i][j]] = true
				black[i/3*3+j/3][board[i][j]-'1'] = true
			}
			if board[j][i] != '.' {
				cols[board[j][i]] = true
			}
		}
	}
	return true
}

func IsValidSudoku(board [][]byte) bool {
	rows, cols, black := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]bool, 9)
		cols[i] = make([]bool, 9)
		black[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				blackIdx := i/3*3 + j/3
				if rows[i][num] || cols[j][num] || black[blackIdx][num] {
					return false
				}
				rows[i][num] = true
				cols[j][num] = true
				black[blackIdx][num] = true
			}
		}
	}
	return true
}
