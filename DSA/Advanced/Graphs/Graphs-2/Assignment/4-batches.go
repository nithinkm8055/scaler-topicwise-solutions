package Assignment

func batches(A int, B []int, C [][]int, D int) int {

	// A is the total number of students
	// C is a M*2 matrix where C[i][0] knows C[i][1]
	// D is the minimum batch strength required for selection

	// use Disjoint set union to create a parent slice
	// find all unique parents, and set of the group

	parent := make([]int, A+1)
	for i := range parent {
		parent[i] = i
	}

	// populating parents
	for i := range C {

		v1 := C[i][0]
		v2 := C[i][1]

		rootOne := _find(v1, parent)
		rootTwo := _find(v2, parent)

		if rootOne != rootTwo {
			parent[rootOne] = rootTwo
		}
	}

	// find unique parents to get strength
	visited := make([]bool, A+1)
	strength := make(map[int]int)
	counter := 0
	for i := 1; i < len(parent); i++ {
		if i != parent[i] {
			parent[i] = _find(parent[i], parent)
		}
		strength[parent[i]] = strength[parent[i]] + B[i-1]

		if !visited[parent[i]] && strength[parent[i]] >= D {
			visited[parent[i]] = true
			counter++
		}
	}

	return counter
}

func _find(v int, parent []int) int {
	for v != parent[v] {
		v = parent[v]
	}
	return v
}
