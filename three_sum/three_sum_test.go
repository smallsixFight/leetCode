package three_sum

import "testing"

func TestTreeSum(t *testing.T) {
	t.Log(TreeSum([]int{-1, 0, 1, 2, -1, -4}))
	t.Log(TreeSum([]int{-1, 0, 1, 0}))
	t.Log(TreeSum([]int{3, 0, -2, -1, 1, 2}))
}
