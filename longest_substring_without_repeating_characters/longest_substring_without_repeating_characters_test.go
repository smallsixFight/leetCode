package longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLengestSubstring(t *testing.T) {
	t.Log(LengthOfLengestSubstring("abcabcbb"))
	t.Log(1 == LengthOfLengestSubstring("bbbbb"))
	t.Log(3 == LengthOfLengestSubstring("pwwkew"))
	t.Log(2 == LengthOfLengestSubstring("abba"))
}
