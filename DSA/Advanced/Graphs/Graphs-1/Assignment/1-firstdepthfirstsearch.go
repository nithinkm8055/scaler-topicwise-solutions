package Assignment

//You are given N towns (1 to N). All towns are connected via unique directed path as mentioned in the input.
//
//Given 2 towns find whether you can reach the first town from the second without repeating any edge.
//
//B C : query to find whether B is reachable from C.
//
//Input contains an integer array A of size N and 2 integers B and C ( 1 <= B, C <= N ).
//
//There exist a directed edge from A[i] to i+1 for every 1 <= i < N. Also, it's guaranteed that A[i] <= i for every 1 <= i < N.
//
//NOTE: Array A is 0-indexed. A[0] = 1 which you can ignore as it doesn't represent any edge.

func solve2(A []int, dest int, src int) int {

	// create adjacency matrix
	graph := _adjList(A)
	visited := make([]int, len(A)+1)
	_dfs(src, graph, visited)

	if visited[dest] == 1 {
		return 1
	}
	return 0
}

func _dfs(src int, graph [][]int, visited []int) {

	visited[src] = 1
	for i := 0; i < len(graph[src]); i++ {
		if visited[graph[src][i]] == 0 {
			_dfs(graph[src][i], graph, visited)
		}
	}
}

func _adjList(A []int) [][]int {
	graph := make([][]int, len(A)+1)
	for i := 1; i < len(A); i++ {
		src := A[i]
		dest := i + 1

		graph[src] = append(graph[src], dest)
	}

	return graph
}
