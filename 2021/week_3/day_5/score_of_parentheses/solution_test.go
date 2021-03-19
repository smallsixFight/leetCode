package score_of_parentheses

import "testing"

func Test_scoreOfParentheses(t *testing.T) {
	t.Log(scoreOfParentheses("()") == 1)
	t.Log(scoreOfParentheses("(())") == 2)
	t.Log(scoreOfParentheses("()()") == 2)
	t.Log(scoreOfParentheses("(()(()))") == 6)
	t.Log(scoreOfParentheses("(())()") == 3)
	t.Log(scoreOfParentheses("((())())") == 6)
	t.Log(scoreOfParentheses("(()())()"))
	t.Log(scoreOfParentheses("((()())())"))
	// 6
}
