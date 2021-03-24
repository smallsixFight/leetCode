package check_if_word_is_valid_after_substitutions

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	t.Log(isValid("aabcbc"))
	t.Log(isValid("abcabcababcc"))
	t.Log(isValid("abccba"))
	t.Log(isValid("cababc"))
}
