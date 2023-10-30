package Assignment

import "sort"

// TODO: Use heap to store result
func solve(A int, B [][]int) []int {
	// A is number of nodes
	// B is directed edge from B[i][0] to B[i][1]

	result := make([]int, 0)

	// adjacency list for nodes
	graph := adjList(A, B)

	// create an indegree array
	indegree := make([]int, A+1)

	for i := 0; i < len(B); i++ {
		dest := B[i][1]
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

		return []int{}
	}

	return final

}

func adjList(A int, B [][]int) [][]int {
	graph := make([][]int, A+1)

	for i := range B {
		src := B[i][0]
		dest := B[i][1]

		graph[src] = append(graph[src], dest)
	}

	return graph

}
