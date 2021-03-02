package valid_parentheses

import "testing"

func Test_isValid(t *testing.T) {
	t.Log(isValid("([)]"))
	t.Log(isValid("{[]}"))
}
