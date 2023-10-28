package Assignment

//Q: Cycle in Directed Graph

//Given an directed graph having A nodes. A matrix B of size M x 2 is given which represents the M edges such that there is a edge directed from node B[i][0] to node B[i][1].
//
//Find whether the graph contains a cycle or not, return 1 if cycle is present else return 0.
//
//NOTE:
//
//The cycle must contain atleast two nodes.
//There are no self-loops in the graph.
//There are no multiple edges between two nodes.
//The graph may or may not be connected.
//Nodes are numbered from 1 to A.
//Your solution will run on multiple test cases. If you are using global variables make sure to clear them.

func _solve(A int, B [][]int) int {
	graph := __adjList(A, B)

	for i := 1; i < A; i++ {
		visited := make([]int, A+1)
		if !__dfs(i, visited, graph) {
			return 1
		}
	}

	return 0
}

func __dfs(src int, visited []int, graph [][]int) bool {
	visited[src] = 1
	for i := 0; i < len(graph[src]); i++ {
		if visited[graph[src][i]] == 0 {
			if !__dfs(graph[src][i], visited, graph) {
				return false
			}
		} else {
			return false
		}
	}
	visited[src] = 0
	return true
}

func __adjList(A int, B [][]int) [][]int {
	graph := make([][]int, A+1)

	for i := 0; i < len(B); i++ {
		src := B[i][0]
		dest := B[i][1]

		graph[src] = append(graph[src], dest)
	}
	return graph
}
