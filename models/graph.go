package models

import "container/heap"

// 拓扑排序 210. Course Schedule II
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	ind := make([]int, numCourses)
	for _, preq := range prerequisites {
		graph[preq[1]] = append(graph[preq[1]], preq[0])
		ind[preq[0]]++
	}
	queue := []int{}
	for i := 0; i < numCourses; i++ {
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

// 迪杰斯特拉算法
func networkDelayTime(times [][]int, n int, k int) []int {
	graph := make([][][2]int, n+1, n+1)
	en := len(times)
	for i := 0; i < en; i++ {
		u, v, w := times[i][0], times[i][1], times[i][2]
		graph[u] = append(graph[u], [2]int{v, w})
	}
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = 600001
	}
	dist[k] = 0

	hp := &Heap2{[2]int{0, k}}
	heap.Init(hp)
	for hp.Len() > 0 {
		edge := heap.Pop(hp).([2]int)
		d, n := edge[0], edge[1]
		if dist[n] < d {
			continue
		}

		for _, nedge := range graph[n] {
			if d := dist[n] + nedge[1]; d < dist[nedge[0]] {
				dist[nedge[0]] = d
				heap.Push(hp, [2]int{d, nedge[0]})
			}

		}
	}
	return dist
}
