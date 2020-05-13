package excel_sheet_column_title

import "testing"

func TestConvertToTitle(t *testing.T) {
	t.Log(ConvertToTitle(52))
	t.Log(ConvertToTitle(701))
	t.Log("A" == ConvertToTitle(1))
	t.Log("AB" == ConvertToTitle(28))
	t.Log("AZ" == ConvertToTitle(52))
	t.Log("ZY" == ConvertToTitle(701))
}
