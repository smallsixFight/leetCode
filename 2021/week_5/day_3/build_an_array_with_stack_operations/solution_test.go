package build_an_array_with_stack_operations

import "testing"

func Test_buildArray(t *testing.T) {
	t.Log(buildArray([]int{1, 3}, 3))
	t.Log(buildArray([]int{1, 2, 3}, 3))
	t.Log(buildArray([]int{1, 2}, 4))
	t.Log(buildArray([]int{2, 3, 4}, 4))
	t.Log(buildArray([]int{3, 4}, 4))
}
