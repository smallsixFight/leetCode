package next_greater_element_i

import "testing"

func Test_nextGreaterElement(t *testing.T) {
	t.Log(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	t.Log(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
