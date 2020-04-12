package string_matching_in_an_array

import "testing"

func TestStringMatching(t *testing.T) {
	t.Log(StringMatching([]string{"mass", "as", "hero", "superhero"}))
	t.Log(StringMatching([]string{"leetcode", "et", "code"}))
	t.Log(StringMatching([]string{"blue", "green", "bu"}))
	t.Log(StringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
}
