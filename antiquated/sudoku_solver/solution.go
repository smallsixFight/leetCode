package sudoku_solver

var (
	rows, cols, boxes [9][10]int
	isSolve           bool
	bo                [][]byte
)

func SolveSuDoku(board [][]byte) {
	rows, cols, boxes = [9][10]int{}, [9][10]int{}, [9][10]int{}
	isSolve = false
	bo = board
	// 初始化记录
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				placeNumber(int(board[i][j]-'0'), i, j)
			}
		}
	}
	backTrack(0, 0)
}

func backTrack(row, col int) {
	if bo[row][col] == '.' {
		for i := 1; i < 10; i++ {
			if couldPlace(i, row, col) {
				placeNumber(i, row, col)
				placeNextNumber(row, col)
				if !isSolve {
					removeNumber(i, row, col)
				}
			}
		}
	} else {
		placeNextNumber(row, col)
	}
}

func placeNextNumber(row, col int) {
	if row == 8 && col == 8 {
		isSolve = true
	} else {
		if col == 8 {
			backTrack(row+1, 0)
		} else {
			backTrack(row, col+1)
		}
	}
}

func couldPlace(num int, row, col int) bool {
	idx := (row/3)*3 + col/3
	return rows[row][num]+cols[col][num]+boxes[idx][num] == 0
}

func placeNumber(num int, row, col int) {
	idx := (row/3)*3 + col/3
	rows[row][num]++
	cols[col][num]++
	boxes[idx][num]++
	bo[row][col] = byte(num + '0')
}

func removeNumber(num int, row, col int) {
	idx := (row/3)*3 + col/3
	rows[row][num]--
	cols[col][num]--
	boxes[idx][num]--
	bo[row][col] = '.'
}
