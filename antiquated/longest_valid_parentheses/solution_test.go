package longest_valid_parentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	t.Log(2 == LongestValidParentheses("(()"))
	t.Log(4 == LongestValidParentheses(")()())"))
	t.Log(8 == LongestValidParentheses("(()((())()()"))
	t.Log(8 == LongestValidParentheses("((((())))(()("))
	t.Log(6 == LongestValidParentheses("(()())"))
}
