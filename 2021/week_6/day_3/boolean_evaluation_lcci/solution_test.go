package boolean_evaluation_lcci

import "testing"

func Test_countEval(t *testing.T) {
	//t.Log(1&1)
	//t.Log(1&0)
	//t.Log(0&0)
	//t.Log(1|1)
	//t.Log(1|0)
	//t.Log(0|0)
	//t.Log(1^1)
	//t.Log(1^0)
	//t.Log(0^0)
	t.Log(countEval("1^0|0|1", 0))
	t.Log(countEval("0&0&0&1^1|0", 1))
}
