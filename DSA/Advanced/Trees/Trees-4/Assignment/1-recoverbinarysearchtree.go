package Assignment

import "sort"

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func recoverTree(A *treeNode) []int {

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
	x := -1
	y := -1

	sorted := make([]int, 0)
	sorted = append(sorted, result...)
	sort.Ints(sorted)

	for i := range result {
		if result[i] != sorted[i] {
			if x == -1 {
				x = sorted[i]
			} else {
				y = sorted[i]
			}
		}
	}

	return []int{x, y}
}
