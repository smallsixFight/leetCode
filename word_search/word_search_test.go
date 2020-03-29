package word_search

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	t.Log(Exist(board, "ABCCED"))
	t.Log(Exist(board, "SEE"))
	t.Log(Exist(board, "ABCB"))
}
