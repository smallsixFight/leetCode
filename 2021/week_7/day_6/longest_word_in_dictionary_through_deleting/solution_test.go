package longest_word_in_dictionary_through_deleting

import "testing"

func Test_findLongestWord(t *testing.T) {
	t.Log(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	t.Log(findLongestWord("abpcplea", []string{"a", "b", "c", "d"}))
}
