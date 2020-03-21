package spiral_matrix

import "testing"

func TestSpiralOrder(t *testing.T) {
	res := SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	t.Log(res)
	res = SpiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
	t.Log(res)
	res = SpiralOrder([][]int{{1, 2}, {3, 4}})
	t.Log(res)
	res = SpiralOrder([][]int{{1, 2, 3, 4, 5, 6}})
	t.Log(res)
	res = SpiralOrder([][]int{{1}, {2}, {3}, {4}, {5}, {6}})
	t.Log(res)
	res = SpiralOrder([][]int{})
	t.Log(res)
	res1 := SpiralOrder([][]int{{2, 5, 8}, {4, 0, 1}})
	t.Log(res1)
}
