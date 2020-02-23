package search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {
	t.Log(SearchInsert([]int{1, 3, 5, 6}, 5))
	t.Log(SearchInsert([]int{1, 3, 5, 6}, 2))
	t.Log(SearchInsert([]int{1, 3, 5, 6}, 7))
	t.Log(SearchInsert([]int{1, 3, 5, 6}, 0))
}
