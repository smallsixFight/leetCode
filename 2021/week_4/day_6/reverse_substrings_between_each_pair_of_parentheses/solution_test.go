package reverse_substrings_between_each_pair_of_parentheses

import "testing"

func Test_reverseParentheses(t *testing.T) {
	t.Log(reverseParentheses("(abcd)"))
	t.Log(reverseParentheses("(u(love)i)"))
	t.Log(reverseParentheses("(ed(et(oc))el)"))
}
