package longest_common_prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	t.Log(LongestCommonPrefix([]string{"abcd", "abccc", "abttt"}))
	t.Log(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	t.Log(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
}
