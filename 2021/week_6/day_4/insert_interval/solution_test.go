package insert_interval

import "testing"

func Test_insert(t *testing.T) {
	t.Log(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})) // [[1 5] [6 9]]
	// [[1 2] [3 10] [12 16]]
	t.Log(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	t.Log(insert([][]int{}, []int{5, 7}))       // [[5 7]]
	t.Log(insert([][]int{{1, 5}}, []int{2, 3})) // [[1 5]]
	t.Log(insert([][]int{{1, 5}}, []int{0, 0})) // [[0 0] [1 5]]
}
