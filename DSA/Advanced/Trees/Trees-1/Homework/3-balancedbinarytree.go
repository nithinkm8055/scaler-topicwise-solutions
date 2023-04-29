package Homework

func isBalanced(A *treeNode) int {
	if A == nil {
		return 1
	}

	x := height(A.left)
	y := height(A.right)

	if abs(x-y) <= 1 && isBalanced(A.left) == 1 && isBalanced(A.right) == 1 {
		return 1
	}

	return 0
}

func height(A *treeNode) int {
	if A == nil {
		return 0
	}

	return 1 + max(height(A.left), height(A.right))
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
