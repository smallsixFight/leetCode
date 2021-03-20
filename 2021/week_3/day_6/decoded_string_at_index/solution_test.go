package decoded_string_at_index

import "testing"

func Test_decodeAtIndex(t *testing.T) {
	t.Log(decodeAtIndex("leet2code3", 10))
	t.Log(decodeAtIndex("ha22", 5))
	t.Log(decodeAtIndex("a2345678999999999999999", 1))
	t.Log(decodeAtIndex("abc", 1))
	t.Log(decodeAtIndex("y959q969u3hb22odq595", 222280369))
}

func Test_decodeAtIndex2(t *testing.T) {
	t.Log(decodeAtIndex2("leet2code3", 10))
	t.Log(decodeAtIndex2("ha22", 5))
	t.Log(decodeAtIndex2("a2345678999999999999999", 1))
	t.Log(decodeAtIndex2("abc", 1))
	t.Log(decodeAtIndex2("y959q969u3hb22odq595", 222280369))
}
