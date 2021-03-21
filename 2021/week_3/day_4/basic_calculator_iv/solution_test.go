package basic_calculator_iv

import (
	"testing"
)

func Test_basicCalculatorIV(t *testing.T) {
	t.Log(basicCalculatorIV("e + 8 - a + 5", []string{"e"}, []int{1}))
	t.Log(basicCalculatorIV("e - 8 + temperature - pressure", []string{"e", "temperature"}, []int{1, 0}))
	t.Log(basicCalculatorIV("e - 8 + temperature * pressure", []string{"e", "temperature"}, []int{1, 12}))
	t.Log(basicCalculatorIV("e - 8 + temperature * temperature * pressure", []string{"e", "temperature"}, []int{1, 12}))
	t.Log(basicCalculatorIV("e - 8 + pressure * temperature", []string{"e", "temperature"}, []int{1, 12}))
	t.Log(basicCalculatorIV("(e + 8) * (e - 8)", []string{}, []int{}))
	t.Log(basicCalculatorIV("7 - 7", []string{}, []int{}))
	t.Log(basicCalculatorIV("a * b * c + b * a * c * 4", []string{}, []int{}))
	t.Log(basicCalculatorIV("((a - b) * (b - c) + (c - a)) * ((a - b) + (b - c) * (c - a))", []string{}, []int{}))
}

func Test_cal(t *testing.T) {
	t.Log(string(cal([]byte("a-b"), []byte("b-c"), '-')))
}

func Test_generateExp(t *testing.T) {
	empty := map[string]int{}
	t.Log(string(generateExp([]byte("e + 8 - a + 5"), map[string]int{"e": 1})))
	t.Log(string(generateExp([]byte("e - 8 + temperature - pressure"), map[string]int{"e": 1, "temperature": 0})))
	t.Log(string(generateExp([]byte("e - 8 + temperature * pressure"), map[string]int{"e": 1, "temperature": 12})))
	t.Log(string(generateExp([]byte("e - 8 + temperature * temperature * pressure"), map[string]int{"e": 1, "temperature": 12})))
	t.Log(string(generateExp([]byte("e - 8 + pressure * temperature"), map[string]int{"e": 1, "temperature": 12})))
	t.Log(string(generateExp([]byte("(e + 8) * (e - 8)"), empty)))
	t.Log(string(generateExp([]byte("7 - 7"), empty)))
	t.Log(string(generateExp([]byte("a * b * c + b * a * c * 4"), empty)))
	t.Log(string(generateExp([]byte("((a - b) * (b - c) + (c - a)) * ((a - b) + (b - c) * (c - a))"), empty)))
}
