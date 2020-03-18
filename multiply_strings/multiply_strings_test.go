package multiply_strings

import "testing"

func TestMultiplyStrings(t *testing.T) {
	t.Log(MultiplyStrings("11", "11"))
	t.Log(MultiplyStrings("2", "3"))
	t.Log(MultiplyStrings("123", "456"))
	t.Log(MultiplyStrings("9", "9"))
	t.Log(MultiplyStrings("12222222222222222222", "211111111111111111111111111111"))
	t.Log(MultiplyStrings("88888888777777777777799999999", "55555555556666999999990000888"))
	t.Log(MultiplyStrings("123456789", "987654321"))
}
