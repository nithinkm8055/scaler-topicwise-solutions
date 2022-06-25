package Assignment

func isSameTree(A *treeNode, B *treeNode) int {

	if A == nil && B == nil {
		return 1
	}

	if A != nil && B != nil {

		l := isSameTree(A.left, B.left)
		r := isSameTree(A.right, B.right)

		if A.value == B.value && l == 1 && r == 1 {
			return 1
		}
		return 0
	}

	return 0

}
