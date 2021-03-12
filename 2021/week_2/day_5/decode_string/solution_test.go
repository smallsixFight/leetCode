package decode_string

import "testing"

func Test_decodeString(t *testing.T) {
	t.Log(decodeString("3[a]2[bc]") == "aaabcbc")
	t.Log(decodeString("3[a2[c]]") == "accaccacc")
	t.Log(decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef")
	t.Log(decodeString("abc3[cd]xyz") == "abccdcdcdxyz")
	t.Log(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef") == "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef")
}
