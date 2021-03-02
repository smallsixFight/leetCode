package implement_strstr

import "testing"

func TestStrstr(t *testing.T) {
	t.Log(2 == Strstr("hello", "ll"))
	t.Log(-1 == Strstr("aaaaa", "bba"))
	t.Log(0 == Strstr("hello", ""))
}

func TestStrstr2(t *testing.T) {
	t.Log(2 == Strstr2("hello", "ll"))
	t.Log(-1 == Strstr2("aaaaa", "bba"))
	t.Log(0 == Strstr2("hello", ""))
}
