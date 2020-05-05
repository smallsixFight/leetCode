package course_schedule

func CanFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 || len(prerequisites) == 0 {
		return true
	}
	adjacency := make(map[int][]int)
	flags := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		adjacency[prerequisites[i][1]] = append(adjacency[prerequisites[i][1]], prerequisites[i][0])
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(adjacency, flags, i) {
			return false
		}
	}
	return true
}

func dfs(adjacency map[int][]int, flags []int, course int) bool {
	if flags[course] == 1 {
		return false
	}
	if flags[course] == -1 {
		return true
	}
	flags[course] = 1
	for i := range adjacency[course] {
		if !dfs(adjacency, flags, adjacency[course][i]) {
			return false
		}
	}
	flags[course] = -1
	return true
}
