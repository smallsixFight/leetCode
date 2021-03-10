package remove_duplicate_letters

import "testing"

func Test_removeDuplicateLetters(t *testing.T) {
	t.Log(removeDuplicateLetters("aecdbabdce") == "abdce")
	t.Log(removeDuplicateLetters("abcdeaedcb") == "abcde")
	t.Log(removeDuplicateLetters("bcdeaedcb") == "aedcb")
	t.Log(removeDuplicateLetters("bcabc") == "abc")
	t.Log(removeDuplicateLetters("cbacdcbc") == "acdb")
	t.Log(removeDuplicateLetters("leetcode"))
	t.Log(removeDuplicateLetters("caccabad"))
}
