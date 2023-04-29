package Assignment

func solve(A *treeNode, B int, C int) int {

	counter := 0
	if B > C {
		B, C = C, B
	}

	var preorder func(*treeNode, int, int)

	preorder = func(A *treeNode, B, C int) {
		if A == nil {
			return
		}
		if A.value >= B && A.value <= C {
			counter++
		}
		preorder(A.left, B, C)
		preorder(A.right, B, C)
	}

	preorder(A, B, C)

	return counter
}
