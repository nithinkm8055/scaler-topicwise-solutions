package Assignment

// Detect cycle in undirected graph
func __solve(A int, B [][]int) int {

	// use disjoint set union to detect cycles

	// initialize parent, mark parent of each vertice as the vertice itself
	parent := make([]int, A+1)
	for i := range parent {
		parent[i] = i
	}

	//
	for i := range B {
		rootOne := find(B[i][0], parent)
		rootTwo := find(B[i][1], parent)

		if rootOne != rootTwo {
			parent[rootOne] = rootTwo
		} else {
			if len(B) > 2 {
				return 1
			}
		}
	}

	return 0

}

func find(v int, parent []int) int {
	if parent[v] == v {
		return v
	}

	return find(parent[v], parent)
}
