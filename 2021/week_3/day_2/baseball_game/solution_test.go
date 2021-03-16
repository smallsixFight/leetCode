package baseball_game

import "testing"

func Test_calPoints(t *testing.T) {
	t.Log(calPoints([]string{"5", "2", "C", "D", "+"}))
	t.Log(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
	t.Log(calPoints([]string{"1"}))
}
