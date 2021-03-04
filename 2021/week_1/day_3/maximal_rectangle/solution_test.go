package maximal_rectangle

import "testing"

func Test_maximalRectangle(t *testing.T) {
	t.Log(maximalRectangle([][]byte{}) == 0)
	t.Log(maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}) == 6)

	t.Log(maximalRectangle([][]byte{{'0'}}) == 0)
	t.Log(maximalRectangle([][]byte{{'1'}}) == 1)
	t.Log(maximalRectangle([][]byte{{'0', '0'}}) == 0)
}
