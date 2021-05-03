package employee_importance

import "testing"

func Test_getImportance(t *testing.T) {
	list := make([]*Employee, 0)
	list = append(list, &Employee{
		Id:           1,
		Importance:   5,
		Subordinates: []int{2, 3},
	})
	list = append(list, &Employee{
		Id:           2,
		Importance:   3,
		Subordinates: []int{},
	})
	list = append(list, &Employee{
		Id:           3,
		Importance:   3,
		Subordinates: []int{},
	})
	t.Log(getImportance(list, 1))
}
