package smallest_subsequence_of_distinct_characters

import "testing"

func Test_smallestSubsequence(t *testing.T) {
	t.Log(smallestSubsequence("aecdbabdce") == "abdce")
	t.Log(smallestSubsequence("abcdeaedcb") == "abcde")
	t.Log(smallestSubsequence("bcdeaedcb") == "aedcb")
	t.Log(smallestSubsequence("bcabc") == "abc")
	t.Log(smallestSubsequence("cbacdcbc") == "acdb")
	t.Log(smallestSubsequence("leetcode"))
	t.Log(smallestSubsequence("caccabad"))
	t.Log(smallestSubsequence("bccddeacacaceabcbbbbdbebbbaaaa"))
}
