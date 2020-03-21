package pow_n

import (
	"math"
	"testing"
)

func TestMyPowTwo(t *testing.T) {
	t.Log(MyPowTwo(2, 10))
	t.Log(MyPowTwo(2, math.MinInt32))
}

func TestMyPow(t *testing.T) {
	t.Log(MyPow(2, 10))
	t.Log(MyPow(2, math.MinInt32))
}
