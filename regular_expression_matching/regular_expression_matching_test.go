package regular_expression_matching

import "testing"

func TestIsMatch2(t *testing.T) {
	t.Log(false == IsMatch2("aa", "a"))
	t.Log(true == IsMatch2("aa", "a*"))
	t.Log(true == IsMatch2("ab", ".*"))
	t.Log(true == IsMatch2("aab", "c*a*b"))
	t.Log(false == IsMatch2("mississippi", "mis*is*p*."))
	t.Log(true == IsMatch2("", ""))
	t.Log(false == IsMatch2("a", ""))
	t.Log(false == IsMatch2("", "a"))
	t.Log(true == IsMatch2("aaa", "a*a"))
	t.Log(true == IsMatch2("aaa", "a*aa"))
	t.Log(true == IsMatch2("aaa", "a*aaa"))
	t.Log(false == IsMatch2("aaa", "a*aaaa"))
	t.Log(true == IsMatch2("aaa", "ab*a*c*a"))
	t.Log(false == IsMatch2("aaa", "aaaa"))
}

// "aaa"
// "a*a"
func TestIsMatch(t *testing.T) {
	t.Log(false == IsMatch("aa", "a"))
	t.Log(true == IsMatch("aa", "a*"))
	t.Log(true == IsMatch("ab", ".*"))
	t.Log(true == IsMatch("aab", "c*a*b"))
	t.Log(false == IsMatch("mississippi", "mis*is*p*."))
	t.Log(true == IsMatch("", ""))
	t.Log(false == IsMatch("a", ""))
	t.Log(false == IsMatch("", "a"))
	t.Log(true == IsMatch("aaa", "a*a"))
	t.Log(true == IsMatch("aaa", "a*aa"))
	t.Log(true == IsMatch("aaa", "a*aaa"))
	t.Log(false == IsMatch("aaa", "a*aaaa"))
	t.Log(true == IsMatch("aaa", "ab*a*c*a"))
	t.Log(false == IsMatch("aaa", "aaaa"))
}
