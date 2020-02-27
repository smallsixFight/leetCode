package longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Log(LongestPalindrome("babad"))
	t.Log(LongestPalindrome("cbbd"))
}

func TestLongestPalindrome2(t *testing.T) {
	t.Log(LongestPalindrome2("babad"))
	t.Log(LongestPalindrome2("cbbd"))
}
