package Assignment

//Given an directed graph having A nodes labelled from 1 to A containing M edges given by matrix B of size M x 2such that there is a edge directed from node
//
//B[i][0] to node B[i][1].
//
//Find whether a path exists from node 1 to node A.
//
//Return 1 if path exists else return 0.
//
//NOTE:
//
//There are no self-loops in the graph.
//There are no multiple edges between two nodes.
//The graph may or may not be connected.
//Nodes are numbered from 1 to A.
//Your solution will run on multiple test cases. If you are using global variables make sure to clear them.

// A, number of nodes in graph
// B directed graph containing 2 edges
func solve(A int, B [][]int) int {

	// create an adjacency matrix
	// traverse adjacency matrix

	adjList := adjList(A, B)
	visited := make([]int, len(adjList)+1)
	dfs(1, adjList, visited) // using dfs for traversal

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

func adjList(A int, B [][]int) [][]int {

	graph := make([][]int, A+1)

	for i := 0; i < len(B); i++ {
		src := B[i][0]
		dest := B[i][1]

		graph[src] = append(graph[src], dest)
	}

	return graph
}
