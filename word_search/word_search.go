package word_search

var words string

func Exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(board)*len(board[0]) < len(word) {
		return false
	}
	words = word
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, 0, map[[2]int]bool{}, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, k int, record map[[2]int]bool, i, j int) bool {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
		return false
	}
	if record[[2]int{i, j}] || board[i][j] != words[k] {
		return false
	}
	record[[2]int{i, j}] = true
	if k == len(words)-1 {
		return true
	}
	if dfs(board, k+1, record, i-1, j) || dfs(board, k+1, record, i+1, j) ||
		dfs(board, k+1, record, i, j-1) || dfs(board, k+1, record, i, j+1) {
		return true
	}
	record[[2]int{i, j}] = false // 当前节点探索完不能返回 true，则将当前节点设置为未使用，以便下一次使用
	return false
}
