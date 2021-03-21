package online_stock_span

import "testing"

func TestConstructor(t *testing.T) {
	s := Constructor()
	t.Log(s.Next(100))
	t.Log(s.Next(80))
	t.Log(s.Next(60))
	t.Log(s.Next(70))
	t.Log(s.Next(60))
	t.Log(s.Next(75))
	t.Log(s.Next(85))
}
