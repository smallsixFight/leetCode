package sort_colors

import "testing"

func TestSortColors(t *testing.T) {
	arr := []int{2, 0, 2, 1, 1, 0}
	SortColors(arr)
	t.Log(arr)
	arr = []int{2, 2, 1, 1}
	SortColors(arr)
	t.Log(arr)
	arr = []int{2, 0, 2, 0, 0, 0}
	SortColors(arr)
	t.Log(arr)
	arr = []int{1, 1, 0, 1, 0, 0, 1}
	SortColors(arr)
	t.Log(arr)
	arr = []int{1, 1, 2, 1, 2, 2, 1}
	SortColors(arr)
	t.Log(arr)
	arr = []int{1, 2, 0}
	SortColors(arr)
	t.Log(arr)
	arr = []int{0, 0}
	SortColors(arr)
	t.Log(arr)
	arr = []int{2, 0, 1}
	SortColors(arr)
	t.Log(arr)
}
