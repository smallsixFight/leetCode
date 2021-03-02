package set_matrix_zeroes

import "testing"

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	SetZeroes(matrix)
	t.Log(matrix)
	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	SetZeroes(matrix)
	t.Log(matrix)
	matrix = [][]int{{1, 1, 1}, {0, 1, 2}}
	SetZeroes(matrix)
	t.Log(matrix)
	matrix = [][]int{{1, 0, 3}}
	SetZeroes(matrix)
	t.Log(matrix)
}
