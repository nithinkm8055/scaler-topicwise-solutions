package Assignment

// A, number of nodes in graph
// B directed graph containing 2 edges
func solve(A int, B [][]int) int {

	// create an adjacency matrix
	// traverse adjacency matrix

	adjMatrix := adjMatrix(A, B)
	visited := make([]int, len(adjMatrix)+1)
	dfs(1, adjMatrix, visited) // using dfs for traversal

	if visited[A] == 1 {
		return 1
	}

	return 0

}

func dfs(source int, graph [][]int, visited []int) {

	// store if a node is visited

	visited[source] = 1

	for i := 0; i < len(graph[source]); i++ {
		if visited[graph[source][i]] == 0 {
			dfs(graph[source][i], graph, visited)
		}
	}

}

func adjMatrix(A int, B [][]int) [][]int {

	graph := make([][]int, A+1)

	for i := 0; i < len(B); i++ {
		src := B[i][0]
		dest := B[i][1]

		graph[src] = append(graph[src], dest)
	}

	return graph
}
