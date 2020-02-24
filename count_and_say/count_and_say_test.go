package count_and_say

import "testing"

func TestCountAndSay(t *testing.T) {
	t.Log("1" == CountAndSay(1))
	t.Log("11" == CountAndSay(2))
	t.Log("21" == CountAndSay(3))
	t.Log("1211" == CountAndSay(4))
	t.Log("111221" == CountAndSay(5))
}

func TestCountAndSay2(t *testing.T) {
	t.Log("1" == CountAndSay2(1))
	t.Log("11" == CountAndSay2(2))
	t.Log("21" == CountAndSay2(3))
	t.Log("1211" == CountAndSay2(4))
	t.Log("111221" == CountAndSay2(5))
}
