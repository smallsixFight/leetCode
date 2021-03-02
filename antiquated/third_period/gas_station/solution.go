package gas_station

func CanCompleteCircuit(gas []int, cost []int) int {
	getSum := func(arr []int) int {
		var sum int
		for i := 0; i < len(arr); i++ {
			sum += arr[i]
		}
		return sum
	}
	if getSum(gas) < getSum(cost) {
		return -1
	}
	curGas := gas[0] - cost[0]
	s := 0
	for i := 1; i < len(gas); i++ {
		if curGas >= 0 {
			curGas += gas[i] - cost[i]
		} else {
			s = i
			curGas = gas[i] - cost[i]
		}
	}
	return s
}
