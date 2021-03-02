package sqrt_x

import "testing"

func TestMySqrt(t *testing.T) {
	t.Log(0 == MySqrt(0))
	t.Log(1 == MySqrt(1))
	t.Log(1 == MySqrt(2))
	t.Log(1 == MySqrt(3))
	t.Log(2 == MySqrt(4))
	t.Log(2 == MySqrt(5))
	t.Log(2 == MySqrt(7))
	t.Log(2 == MySqrt(8))
}

func TestMySqrt2(t *testing.T) {
	t.Log(0 == MySqrt2(0))
	t.Log(1 == MySqrt2(1))
	t.Log(1 == MySqrt2(2))
	t.Log(1 == MySqrt2(3))
	t.Log(2 == MySqrt2(4))
	t.Log(2 == MySqrt2(5))
	t.Log(2 == MySqrt2(7))
	t.Log(2 == MySqrt2(8))
}
