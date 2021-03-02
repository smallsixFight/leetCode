package merge_intervals

import "testing"

func TestMerge(t *testing.T) {
	t.Log(Merge([][]int{{1, 3}, {3, 5}, {2, 4}}))
	t.Log(Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	t.Log(Merge([][]int{{1, 4}, {4, 5}}))
}
