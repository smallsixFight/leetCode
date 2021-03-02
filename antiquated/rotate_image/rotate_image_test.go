package rotate_image

import "testing"

func TestRotate(t *testing.T) {
	matrix := [][]int{{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}
	Rotate(matrix)
	t.Log(matrix)

	matrix2 := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	Rotate(matrix2)
	t.Log(matrix2)

}
