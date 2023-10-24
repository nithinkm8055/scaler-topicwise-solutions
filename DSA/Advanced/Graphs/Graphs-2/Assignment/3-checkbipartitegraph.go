package Assignment

// check if graph is bipartite
// given an undirected graph

func _bipartite(A int, B [][]int) int {

	// create adjacency list
	// traverse through graph (dfs)

	color := make([]int, A)
	for i := range color {
		color[i] = -1
	}

	graph := createAdjList(A, B)
	for i := 0; i < A; i++ {

		if color[i] == -1 {
			color[i] = 0
			if !dfs(i, graph, color) {
				return 0
			}
		}
	}
	return 1
}

func dfs(src int, graph [][]int, color []int) bool {

	for i := 0; i < len(graph[src]); i++ {

		node := graph[src][i]
		if color[node] == -1 {
			color[node] = 1 - color[src]
			if !dfs(node, graph, color) {
				return false
			}
		}
		if color[node] == color[src] {
			return false
		}
	}

	return true
}

func createAdjList(A int, B [][]int) [][]int {

	graph := make([][]int, A)

	for i := 0; i < len(B); i++ {

		src := B[i][0]
		dest := B[i][1]

		graph[src] = append(graph[src], dest)
		graph[dest] = append(graph[dest], src)

	}

	return graph
}
