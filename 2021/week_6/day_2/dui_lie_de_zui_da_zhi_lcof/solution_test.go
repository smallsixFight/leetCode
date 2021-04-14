package dui_lie_de_zui_da_zhi_lcof

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	queue := Constructor()
	queue.Push_back(1)
	queue.Push_back(2)
	t.Log(queue.Max_value())
	t.Log(queue.Pop_front())
	t.Log(queue.Pop_front())

	t.Log("------------------")
	q2 := Constructor()
	t.Log(q2.Pop_front())
	t.Log(q2.Max_value())
}

func TestConstructor2(t *testing.T) {
	v1 := []string{"MaxQueue", "max_value", "pop_front", "max_value", "push_back", "max_value", "pop_front", "max_value", "pop_front", "push_back", "pop_front", "pop_front", "pop_front", "push_back", "pop_front", "max_value", "pop_front", "max_value", "push_back", "push_back", "max_value", "push_back", "max_value", "max_value", "max_value", "push_back", "pop_front", "max_value", "push_back", "max_value", "max_value", "max_value", "pop_front", "push_back", "push_back", "push_back", "push_back", "pop_front", "pop_front", "max_value", "pop_front", "pop_front", "max_value", "push_back", "push_back", "pop_front", "push_back", "push_back", "push_back", "push_back", "pop_front", "max_value", "push_back", "max_value", "max_value", "pop_front", "max_value", "max_value", "max_value", "push_back", "pop_front", "push_back", "pop_front", "max_value", "max_value", "max_value", "push_back", "pop_front", "push_back", "push_back", "push_back", "pop_front", "max_value", "pop_front", "max_value", "max_value", "max_value", "pop_front", "push_back", "pop_front", "push_back", "push_back", "pop_front", "push_back", "pop_front", "push_back", "pop_front", "pop_front", "push_back", "pop_front", "pop_front", "pop_front", "push_back", "push_back", "max_value", "push_back", "pop_front", "push_back", "push_back", "pop_front"}
	v2 := []int{0, 0, 0, 0, 46, 0, 0, 0, 0, 868, 0, 0, 0, 525, 0, 0, 0, 0, 123, 646, 0, 229, 0, 0, 0, 871, 0, 0, 285, 0, 0, 0, 0, 45, 140, 837, 545, 0, 0, 0, 0, 0, 0, 561, 237, 0, 633, 98,
		806, 717, 0, 0, 186, 0, 0, 0, 0, 0, 0, 268, 0, 29, 0, 0, 0, 0, 866, 0, 239, 3, 850, 0, 0, 0, 0, 0, 0, 0, 310, 0, 674, 770, 0, 525, 0, 425, 0, 0, 720, 0, 0, 0, 373, 411, 0, 831, 0, 765, 701, 0}

	q := Constructor()
	for i := 1; i < len(v1); i++ {
		//fmt.Println("no", i, v1[i])
		if v1[i] == "max_value" {
			fmt.Print(q.Max_value(), ",")
		} else if v1[i] == "pop_front" {
			fmt.Print(q.Pop_front(), ",")
		} else {
			q.Push_back(v2[i])
			fmt.Print("null,")
		}
	}
}
