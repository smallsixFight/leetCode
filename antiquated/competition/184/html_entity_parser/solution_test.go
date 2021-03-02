package html_entity_parser

import "testing"

func TestEntityParser(t *testing.T) {
	t.Log(EntityParser("&amp; is an HTML entity but &ambassador; is not."))
	t.Log(EntityParser("and I quote: &quot;...&quot;"))
	t.Log(EntityParser("Stay home! Practice on Leetcode :)"))
	t.Log(EntityParser("x &gt; y &amp;&amp; x &lt; y is always false"))
	t.Log(EntityParser("leetcode.com&frasl;problemset&frasl;all"))
}
