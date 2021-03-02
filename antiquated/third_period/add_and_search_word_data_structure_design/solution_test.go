package add_and_search_word_data_structure_design

import (
	"os"
	"runtime/pprof"
	"testing"
)

// ["WordDictionary","addWord","addWord","search","search","search","search","search","search"]
//[[],["a"],["a"],["."],["a"],["aa"],["a"],[".a"],["a."]]
func TestWordDictionary2(t *testing.T) {
	f, _ := os.Create("pprof.profile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	c := Constructor()
	c.AddWord("a")
	c.AddWord("a")
	t.Log(true == c.Search("."))
	t.Log(true == c.Search("a"))
	t.Log(false == c.Search("aa"))
	t.Log(false == c.Search(".a"))
	t.Log(false == c.Search("a."))
}

func TestWordDictionary(t *testing.T) {
	c := Constructor()
	c.AddWord("bad")
	c.AddWord("dad")
	c.AddWord("mad")
	t.Log(false == c.Search("pad"))
	t.Log(true == c.Search("bad"))
	t.Log(true == c.Search(".ad"))
	t.Log(true == c.Search("b.."))
	t.Log(false == c.Search("badd"))

}
