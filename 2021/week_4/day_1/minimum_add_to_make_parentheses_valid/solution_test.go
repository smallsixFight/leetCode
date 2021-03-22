package minimum_add_to_make_parentheses_valid

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	t.Log(minAddToMakeValid("())"))
	t.Log(minAddToMakeValid("((("))
	t.Log(minAddToMakeValid("()"))
	t.Log(minAddToMakeValid("()))(("))
	t.Log(minAddToMakeValid(""))
}
