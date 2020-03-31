package subsets_ii

import "testing"

func TestSubsetsWithDup(t *testing.T) {
	t.Log(subsetsWithDup([]int{1, 2, 2}))
	t.Log(subsetsWithDup([]int{1, 2, 2, 2}))
	t.Log(subsetsWithDup([]int{1, 2, 2, 3}))
	t.Log(subsetsWithDup([]int{1, 2, 2, 3, 3}))
	t.Log(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}
