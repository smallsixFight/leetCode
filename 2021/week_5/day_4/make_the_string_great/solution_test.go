package make_the_string_great

import "testing"

func Test_makeGood(t *testing.T) {
	t.Log(makeGood("leEeetcode"))
	t.Log(makeGood("abBAcC"))
	t.Log(makeGood("s"))
}
