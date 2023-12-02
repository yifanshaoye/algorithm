package models

// 拓扑排序 210. Course Schedule II
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	ind := make([]int, numCourses)
	for _, preq := range prerequisites {
		graph[preq[1]] = append(graph[preq[1]], preq[0])
		ind[preq[0]]++
	}
	queue := []int{}
	for i:=0; i<numCourses; i++ {
		if ind[i] == 0 {
			queue = append(queue, i)
		}
	}
	res := []int{}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		res = append(res, u)
		for _, v := range graph[u] {
			ind[v]--
			if ind[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return []int{}
}