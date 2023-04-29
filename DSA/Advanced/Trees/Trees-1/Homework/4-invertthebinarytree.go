package Homework

func invertTree(A *treeNode) *treeNode {

	if A == nil {
		return A
	}

	A.left, A.right = A.right, A.left
	invertTree(A.left)
	invertTree(A.right)

	return A

}
