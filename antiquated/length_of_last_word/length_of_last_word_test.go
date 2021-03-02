package length_of_last_word

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	t.Log(1 == LengthOfLastWord("a "))
	t.Log(5 == LengthOfLastWord("Hello World"))
	t.Log(5 == LengthOfLastWord("Hello World  "))
	t.Log(0 == LengthOfLastWord("  "))
}
