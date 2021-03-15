package next_greater_element_ii

import "testing"

func Test_nextGreaterElements(t *testing.T) {
	t.Log(nextGreaterElements([]int{1, 2, 1}))
	t.Log(nextGreaterElements([]int{}))
	t.Log(nextGreaterElements([]int{1, 4, 3, 4, 1}))
	t.Log(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
