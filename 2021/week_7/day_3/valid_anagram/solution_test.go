package sort_list

import "testing"

func Test_isAnagram2(t *testing.T) {
	t.Log(isAnagram2("anagram", "nagaram")) // true
	t.Log(isAnagram2("rat", "car"))         // false
	t.Log(isAnagram2("ab", "a"))            // false
	t.Log(isAnagram2("aa", "a"))            // false
	t.Log(isAnagram2("aacc", "ccac"))       // false
}

func Test_isAnagram(t *testing.T) {
	t.Log(isAnagram("anagram", "nagaram")) // true
	t.Log(isAnagram("rat", "car"))         // false
	t.Log(isAnagram("ab", "a"))            // false
	t.Log(isAnagram("aa", "a"))            // false
	t.Log(isAnagram("aacc", "ccac"))       // false
}
