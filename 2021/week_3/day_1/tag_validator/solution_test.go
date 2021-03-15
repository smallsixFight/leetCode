package tag_validator

import "testing"

func Test_isValid(t *testing.T) {
	t.Log(isValid(""))
	t.Log(isValid("<DIV>This is the first line <![CDATA[<div>]]></DIV>"))
	t.Log(isValid("<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"))

	t.Log(isValid("<A>  <B> </A>   </B>"))
	t.Log(isValid("<DIV>  div tag is not closed  <DIV>"))
	t.Log(isValid("<DIV>  unmatched <  </DIV>"))
	t.Log(isValid("<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"))
	t.Log(isValid("<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>"))
	t.Log(isValid("<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"))
	t.Log(isValid("<A></A><B></B>"))
	t.Log(isValid("<A"))
	t.Log(isValid("<![CDATA[ABC]]><TAG>sometext</TAG>"))
}
