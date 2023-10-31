package Assignment

import "sort"

func _solve(A int, B []int, C []int) int {

	// A is number of nodes
	// B is directed edge from B[i][0] to B[i][1]

	result := make([]int, 0)

	// adjacency list for nodes
	graph := _adjList(A, B, C)

	// create an indegree array
	indegree := make([]int, A+1)

	for i := 0; i < len(B); i++ {
		dest := C[i]
		indegree[dest]++
	}

	// traverse nodes with indegree 0

	for i := 1; i <= A; i++ {
		if indegree[i] == 0 {

			result = append(result, i)

		}
	}

	final := make([]int, 0)

	sort.Ints(result)
	cnt := 0
	for len(result) > 0 {
		sort.Ints(result)
		k := result[0]
		final = append(final, result[0])
		result = result[1:]
		cnt++

		for i := 0; i < len(graph[k]); i++ {
			v := graph[k][i]
			indegree[v]--

			if indegree[v] == 0 {
				result = append(result, v)
			}

		}
	}

	if cnt != A {

		return 0
	}

	return 1

}

func _adjList(A int, B, C []int) [][]int {
	graph := make([][]int, A+1)

	for i := range B {
		src := B[i]
		dest := C[i]

		graph[src] = append(graph[src], dest)
	}

	return graph

}
