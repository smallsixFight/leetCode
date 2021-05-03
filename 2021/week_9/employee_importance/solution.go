package employee_importance

/*
690
来源：[leetCode](https://leetcode-cn.com/)
题目：[员工的重要性](https://leetcode-cn.com/problems/employee-importance/)
标签：`深度优先探索` `广度优先探索` `哈希表`

简单思路：先遍历一遍 `employees`，生成一个对应的 hashTable，然后找到目标 id 的所有直接下属，计算重要度，然后再遍历直接下属的直接下属，以此类推，知道结束，然后累计的重要度结果值。

这里采用 HashTable+BFS 处理。

官网运行结果记录
执行用时：12ms(87%)/16ms(53%)
内存消耗：6.7MB(83%)

*/

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	dict1, dict2 := make(map[int]int, len(employees)), make(map[int][]int, len(employees))
	for i := range employees {
		dict1[employees[i].Id] = employees[i].Importance
		dict2[employees[i].Id] = employees[i].Subordinates
	}
	res, arr := dict1[id], dict2[id]
	for len(arr) > 0 {
		l := len(arr)
		for i := 0; i < l; i++ {
			res += dict1[arr[i]]
			arr = append(arr, dict2[arr[i]]...)
		}
		arr = arr[l:]
	}
	return res
}
