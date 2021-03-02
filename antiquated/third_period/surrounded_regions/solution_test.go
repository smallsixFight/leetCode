package surrounded_regions

import "testing"

func TestSolve(t *testing.T) {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	Solve(board)
	t.Log(board)
}
