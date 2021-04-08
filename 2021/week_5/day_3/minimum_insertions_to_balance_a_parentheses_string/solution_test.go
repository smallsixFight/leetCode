package minimum_insertions_to_balance_a_parentheses_string

import "testing"

func Test_minInsertions(t *testing.T) {
	t.Log(minInsertions("(()))") == 1)
	t.Log(minInsertions("())") == 0)
	t.Log(minInsertions("))())(") == 3)
	t.Log(minInsertions("((((((") == 12)
}
