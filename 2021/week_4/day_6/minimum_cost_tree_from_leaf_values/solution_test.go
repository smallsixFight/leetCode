package minimum_cost_tree_from_leaf_values

import "testing"

func Test_mctFromLeafValues(t *testing.T) {
	t.Log(mctFromLeafValues([]int{6, 2, 4}))
	t.Log(mctFromLeafValues([]int{6, 9, 6, 15, 15}))
	t.Log(mctFromLeafValues([]int{7, 12, 8, 10}))
}

// 84 56 120
