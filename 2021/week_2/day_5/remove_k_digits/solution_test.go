package remove_k_digits

import "testing"

func Test_removeKdigits(t *testing.T) {
	t.Log(removeKdigits("1432219", 3))
	t.Log(removeKdigits("10200", 1))
	t.Log(removeKdigits("10", 2))
	t.Log(removeKdigits("112", 1))
	t.Log(removeKdigits("1111111", 3))
	t.Log(removeKdigits("111111121111112", 3))
	t.Log(removeKdigits("2222111222", 4))
}
