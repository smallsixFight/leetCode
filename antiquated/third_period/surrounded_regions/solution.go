package surrounded_regions

func Solve(board [][]byte) {
	if len(board) < 2 || len(board[0]) < 2 {
		return
	}
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
		if board[len(board)-1][i] == 'O' {
			dfs(board, len(board)-1, i)
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			dfs(board, i, len(board[0])-1)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	board[i][j] = '#'
	if i > 0 && board[i-1][j] == 'O' {
		dfs(board, i-1, j)
	}
	if i < len(board)-1 && board[i+1][j] == 'O' {
		dfs(board, i+1, j)
	}
	if j > 0 && board[i][j-1] == 'O' {
		dfs(board, i, j-1)
	}
	if j < len(board[0])-1 && board[i][j+1] == 'O' {
		dfs(board, i, j+1)
	}
}
