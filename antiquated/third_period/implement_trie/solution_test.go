package implement_trie

import "testing"

func TestTrie2(t *testing.T) {
	c := Constructor()
	c.Insert("app")
	c.Insert("apple")
	c.Insert("beer")
	c.Insert("add")
	c.Insert("jam")
	c.Insert("rental")
	t.Log(c.Search("apps"))
	t.Log(c.Search("app"))
	t.Log(c.Search("ad"))
	t.Log(c.Search("applepie"))
	t.Log(c.Search("rest"))
	t.Log(c.Search("jan"))
	t.Log(c.Search("rent"))
	t.Log(c.Search("beer"))
	t.Log(c.Search("jam"))
	t.Log(c.StartsWith("apps"))
	t.Log(c.StartsWith("app"))
	t.Log(c.StartsWith("ad"))
	t.Log(c.StartsWith("applepie"))
	t.Log(c.StartsWith("rest"))
	t.Log(c.StartsWith("jan"))
	t.Log(c.StartsWith("rent"))
	t.Log(c.StartsWith("beer"))
	t.Log(c.StartsWith("jam"))
}

func TestTrie(t *testing.T) {
	c := Constructor()
	c.Insert("apple")
	t.Log(true == c.Search("apple"))
	t.Log(false == c.Search("app"))
	t.Log(true == c.StartsWith("app"))
	c.Insert("app")
	t.Log(true == c.Search("app"))
}
