package search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	t.Log(0 == Search([]int{1, 2, 3, 4, 5, 6}, 1))
	t.Log(5 == Search([]int{4, 5, 6, 7, 0, 1, 2}, 1))
	t.Log(3 == Search([]int{6, 7, 0, 1, 2, 4, 5}, 1))
	t.Log(0 == Search([]int{1, 2, 4, 5, 6, 7, 0}, 1))
	t.Log(-1 == Search([]int{1, 2, 4, 5, 6, 7, 0}, -1))
	t.Log(Search([]int{3, 1}, 1))
}
