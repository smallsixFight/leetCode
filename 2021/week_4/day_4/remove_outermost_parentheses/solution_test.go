package remove_outermost_parentheses

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
	t.Log(removeOuterParentheses("(()())(())"))
	t.Log(removeOuterParentheses("(()())(())(()(()))"))
	t.Log(removeOuterParentheses("()()"))
	t.Log(removeOuterParentheses("((()())(())(()(())))"))
}
