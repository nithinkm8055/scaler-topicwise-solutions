package Homework

var count int

func solve(A *treeNode) int {
	count = 0
	dfs(A, -1)
	return count

}

func dfs(A *treeNode, maxx int) {
	// Base case
	if A == nil {
		return
	}
	// Increment the count if the current
	// node's value is greater than the
	// maximum value in it's ancestors
	if A.value > maxx {
		count++
	}

	// Left traversal
	dfs(A.left, max(maxx, A.value))

	// Right traversal
	dfs(A.right, max(maxx, A.value))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
