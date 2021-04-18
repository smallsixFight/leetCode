package sort_colors

import "testing"

func Test_removeDuplicates(t *testing.T) {
	arr := []int{1, 1, 2}
	t.Log(arr[:removeDuplicates(arr)])
	arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(arr[:removeDuplicates(arr)])
}
