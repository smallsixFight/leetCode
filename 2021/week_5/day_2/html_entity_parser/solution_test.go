package html_entity_parser

import "testing"

func Test_entityParser2(t *testing.T) {
	t.Log(entityParser2("&amp; is an HTML entity but &ambassador; is not.") == "& is an HTML entity but &ambassador; is not.")
	t.Log(entityParser2("and I quote: &quot;...&quot;") == "and I quote: \"...\"")
	t.Log(entityParser2("Stay home! Practice on Leetcode :)") == "Stay home! Practice on Leetcode :)")
	t.Log(entityParser2("x &gt; y &amp;&amp; x &lt; y is always false") == "x > y && x < y is always false")
	t.Log(entityParser2("leetcode.com&frasl;problemset&frasl;all") == "leetcode.com/problemset/all")
	t.Log(entityParser2("&amp;amp;amp;gt;"))
}

func Test_entityParser(t *testing.T) {
	t.Log(entityParser("&amp; is an HTML entity but &ambassador; is not.") == "& is an HTML entity but &ambassador; is not.")
	t.Log(entityParser("and I quote: &quot;...&quot;") == "and I quote: \"...\"")
	t.Log(entityParser("Stay home! Practice on Leetcode :)") == "Stay home! Practice on Leetcode :)")
	t.Log(entityParser("x &gt; y &amp;&amp; x &lt; y is always false") == "x > y && x < y is always false")
	t.Log(entityParser("leetcode.com&frasl;problemset&frasl;all") == "leetcode.com/problemset/all")
	t.Log(entityParser("&amp;amp;amp;gt;"))
}
