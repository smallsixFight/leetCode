package zigzag_conversion

import "testing"

func TestConvert(t *testing.T) {
	t.Log("LCIRETOESIIGEDHN" == Convert("LEETCODEISHIRING", 3))
	t.Log("LDREOEIIECIHNTSG" == Convert("LEETCODEISHIRING", 4))
}
