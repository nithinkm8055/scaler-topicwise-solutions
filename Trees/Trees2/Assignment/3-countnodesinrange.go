package Assignment

func solve(A *treeNode , B int , C int )  (int) {
	count := 0
	var countValues func(*treeNode)

	countValues = func(A *treeNode) {
		if A == nil {
			return
		}

		if A.value >= B && A.value <= C {
			count++
		}

		countValues(A.left)
		countValues(A.right)
	}

	countValues(A)
	return count
}


