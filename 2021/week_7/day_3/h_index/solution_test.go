package sort_list

import "testing"

func Test_hIndex(t *testing.T) {
	t.Log(hIndex([]int{3, 0, 6, 1, 5}))
	t.Log(hIndex([]int{1, 3, 1}))
	t.Log(hIndex([]int{100}))
}
