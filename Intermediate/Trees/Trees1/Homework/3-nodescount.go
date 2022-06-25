package Homework

func nodeCount(A *treeNode) int {

	if A == nil {
		return 0
	}

	return 1 + nodeCount(A.left) + nodeCount(A.right)
}
