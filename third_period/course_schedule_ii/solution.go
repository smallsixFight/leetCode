package course_schedule_ii

var idx int

func FindOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	}
	adjacency := make(map[int][]int)
	flags := make([]int, numCourses)
	res := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		adjacency[prerequisites[i][1]] = append(adjacency[prerequisites[i][1]], prerequisites[i][0])
	}
	idx = numCourses - 1
	for i := 0; i < numCourses; i++ {
		if !dfs(adjacency, flags, i, res) {
			return []int{}
		}
	}
	return res
}

func dfs(adjacency map[int][]int, flags []int, course int, res []int) bool {
	if flags[course] == 1 {
		return false
	}
	if flags[course] == -1 {
		return true
	}
	flags[course] = 1
	for i := 0; i < len(adjacency[course]); i++ {
		if !dfs(adjacency, flags, adjacency[course][i], res) {
			return false
		}
	}
	res[idx] = course
	idx--
	flags[course] = -1
	return true
}
