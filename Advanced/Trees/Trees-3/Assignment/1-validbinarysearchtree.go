package Assignment

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func isValidBST(A *treeNode) int {

	result := make([]int, 0)

	var inorder func(*treeNode)

	inorder = func(A *treeNode) {
		if A == nil {
			return
		}

		inorder(A.left)
		result = append(result, A.value)
		inorder(A.right)
	}

	inorder(A)

	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			return 0
		}
	}

	return 1
}
