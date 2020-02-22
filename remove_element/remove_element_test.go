package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	t.Log(2 == RemoveElement([]int{3, 2, 2, 3}, 3))
	t.Log(5 == RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	t.Log(6 == RemoveElement([]int{1, 2, 3, 4, 5, 6, 7}, 7))
}
