package three_sum_closest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	t.Log(ThreeSumClosest([]int{-1, 2, 1, -4}, 1))
	t.Log(ThreeSumClosest([]int{1, 1, 1, 1}, 0))
}
