package gas_station

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	t.Log(CanCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	t.Log(CanCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
