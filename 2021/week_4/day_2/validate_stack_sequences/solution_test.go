package validate_stack_sequences

import "testing"

func Test_validateStackSequences(t *testing.T) {
	t.Log(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	t.Log(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
