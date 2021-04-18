package sort_colors

import "testing"

func Test_sortColors(t *testing.T) {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	t.Log(arr)
	arr = []int{2, 0, 1}
	sortColors(arr)
	t.Log(arr)
	arr = []int{0}
	sortColors(arr)
	t.Log(arr)
	arr = []int{1}
	sortColors(arr)
	t.Log(arr)
	arr = []int{1, 2, 0}
	sortColors(arr)
	t.Log(arr)
}
