package group_anagrams

import "testing"

func TestGroupAnagramsThree(t *testing.T) {
	t.Log(GroupAnagramsThree([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(GroupAnagramsThree([]string{"", "eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(GroupAnagramsThree([]string{"tea", "and", "ace", "ad", "eat", "dans"}))
}

func TestGroupAnagramsTwoT(t *testing.T) {
	t.Log(GroupAnagramsTwoT([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(GroupAnagramsTwoT([]string{"", "eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(GroupAnagramsTwoT([]string{"tea", "and", "ace", "ad", "eat", "dans"}))
}

func TestGroupAnagramsTwo(t *testing.T) {
	t.Log(GroupAnagramsTwo([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(GroupAnagramsTwo([]string{"", "eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(GroupAnagramsTwo([]string{"tea", "and", "ace", "ad", "eat", "dans"}))
}

func TestGroupAnagrams(t *testing.T) {
	t.Log(GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(GroupAnagrams([]string{"", " ", "  ", "eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(GroupAnagrams([]string{"tea", "and", "ace", "ad", "eat", "dans"}))
}
