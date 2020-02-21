package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	t.Log(true, "-->", IsValid(""))
	t.Log(true, "-->", IsValid("()"))
	t.Log(true, "-->", IsValid("()[]{}"))
	t.Log(false, "-->", IsValid("(]"))
	t.Log(false, "-->", IsValid("([)]"))
	t.Log(true, "-->", IsValid("{[]}"))
}

func TestIsValid2(t *testing.T) {
	t.Log(true, "-->", IsValid2(""))
	t.Log(true, "-->", IsValid2("()"))
	t.Log(true, "-->", IsValid2("()[]{}"))
	t.Log(false, "-->", IsValid2("(]"))
	t.Log(false, "-->", IsValid2("([)]"))
	t.Log(true, "-->", IsValid2("{[]}"))
}
