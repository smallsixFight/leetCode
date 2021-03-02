package reverse_words_in_a_string

import "testing"

func TestReverseWords(t *testing.T) {
	t.Log(ReverseWords("the sky is blue"))
	t.Log(ReverseWords("  hello world!  "))
	t.Log(ReverseWords("a good   example"))
	t.Log(ReverseWords("   "))
	t.Log(ReverseWords(""))
}
