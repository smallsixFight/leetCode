package minimum_remove_to_make_valid_parentheses

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	t.Log(minRemoveToMakeValid("lee(t(c)o)de)"))
	t.Log(minRemoveToMakeValid("a)b(c)d"))
	t.Log(minRemoveToMakeValid("))(("))
	t.Log(minRemoveToMakeValid("(a(b(c)d)"))
}
