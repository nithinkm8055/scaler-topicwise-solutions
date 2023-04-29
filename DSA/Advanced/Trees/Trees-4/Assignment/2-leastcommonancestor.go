package Assignment

// TODO: refactoring
func lca(A *treeNode, B int, C int) int {

	result1 := make([]*treeNode, 0)
	result2 := make([]*treeNode, 0)

	var preorder1 func(*treeNode, int) bool
	var preorder2 func(*treeNode, int) bool

	preorder1 = func(A *treeNode, search int) bool {

		if A == nil {
			return false
		}

		if A.value == search {
			result1 = append(result1, A)
			return true
		}

		if preorder1(A.left, search) || preorder1(A.right, search) {
			result1 = append(result1, A)
			return true
		}

		return false
	}

	preorder2 = func(A *treeNode, search int) bool {

		if A == nil {
			return false
		}

		if A.value == search {
			result2 = append(result2, A)
			return true
		}

		if preorder2(A.left, search) || preorder2(A.right, search) {
			result2 = append(result2, A)
			return true
		}

		return false
	}

	preorder1(A, B)
	preorder2(A, C)

	if len(result2) < len(result1) {
		result1, result2 = result2, result1
	}
	// fmt.Println(result1, result2)
	for j := 0; j < len(result1); j++ {

		for i := range result2 {
			if result1[j] == result2[i] {
				return result1[j].value
			}
		}

	}
	return -1
}
