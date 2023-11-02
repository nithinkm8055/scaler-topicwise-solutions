package Assignment

import "sort"

func solve(A int, B [][]int) int {
	parent := make([]int, A)
	for i := range parent {
		parent[i] = i
	}

	sort.SliceStable(B, func(i, j int) bool {
		return B[i][2] < B[j][2]
	})

	cost := 0
	for i := range B {
		if union(B[i][0], B[i][1], parent) {
			cost += B[i][2]
		}
	}

	return cost % (1000000007)

}

func find(v int, parent []int) int {
	for v != parent[v] {
		v = parent[v]
	}

	return v
}

func union(x, y int, parent []int) bool {

	rootOne := find(x, parent)
	rootTwo := find(y, parent)

	if rootOne != rootTwo {
		parent[rootOne] = rootTwo
		return true
	}

	return false
}
