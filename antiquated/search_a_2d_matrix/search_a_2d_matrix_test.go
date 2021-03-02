package search_a_2d_matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	t.Log(SearchMatrix(matrix, 3))
	t.Log(SearchMatrix(matrix, 13))
	matrix = [][]int{{1}}
	t.Log(SearchMatrix(matrix, 1))
}
