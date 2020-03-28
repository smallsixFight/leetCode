package subsets

import "testing"

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(Subsets(nums))
	t.Log(Subsets([]int{1, 2}))
	t.Log(Subsets([]int{1, 2, 3, 4}))
}
