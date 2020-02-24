package maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	t.Log(MaxSubArray([]int{1, 2}))
	t.Log(MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
