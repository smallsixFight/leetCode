package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Log(IsPalindrome(" "))
	t.Log(IsPalindrome(",."))
	t.Log(IsPalindrome("ab31ba"))
	t.Log(IsPalindrome("A man, a plan, a canal: Panama"))
	t.Log(IsPalindrome("race a car"))
}
