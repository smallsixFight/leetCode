package decode_ways

import "testing"

func TestNumDecodingsTwo(t *testing.T) {
	s := "0"
	t.Log(0 == NumDecodingsTwo(s))
	s = "1"
	t.Log(1 == NumDecodingsTwo(s))
	s = "01"
	t.Log(0 == NumDecodingsTwo(s))
	s = "10"
	t.Log(1 == NumDecodingsTwo(s))
	s = "12"
	t.Log(2 == NumDecodingsTwo(s))
	s = "226"
	t.Log(3 == NumDecodingsTwo(s))
	s = "301"
	t.Log(0 == NumDecodingsTwo(s))
	s = "1280"
	t.Log(0 == NumDecodingsTwo(s))
	s = "12801"
	t.Log(0 == NumDecodingsTwo(s))
}

func TestNumDecodings(t *testing.T) {
	s := "1"
	t.Log(NumDecodings(s))
	s = "01"
	t.Log(NumDecodings(s))
	s = "12"
	t.Log(NumDecodings(s))
	s = "226"
	t.Log(NumDecodings(s))
	s = "1280"
	t.Log(NumDecodings(s))
	s = "12801"
	t.Log(NumDecodings(s))
}
