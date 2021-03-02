package wildcard_matching

import "testing"

func TestIsMatch(t *testing.T) {
	t.Log(false == IsMatch("aa", "a"))
	t.Log(true == IsMatch("aa", "*"))
	t.Log(false == IsMatch("cb", "?a"))
	t.Log(true == IsMatch("adceb", "*a*b"))
	t.Log(false == IsMatch("acdcb", "a*c?b"))
	t.Log(true == IsMatch("aa", "aa"))
	t.Log(true == IsMatch("aaaa", "***a"))
	t.Log(true == IsMatch("a", "a*"))
	t.Log(true == IsMatch("c", "*?*"))
	t.Log(true == IsMatch("ho", "ho**"))

	t.Log(false == IsMatch("ab", "*a"))
	t.Log(false == IsMatch("aabab", "?**aa?"))
	t.Log(true == IsMatch("aac", "*ac"))

	t.Log(true == IsMatch("abb", "**??"))
	t.Log(false == IsMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
		"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))
}
