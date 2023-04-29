package Assignment

import "sort"

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func kthsmallest(A *treeNode, B int) int {

	var inorder func(*treeNode)
	result := make([]int, 0)
	inorder = func(A *treeNode) {
		if A == nil {
			return
		}

		result = append(result, A.value)
		inorder(A.left)
		inorder(A.right)
	}

	inorder(A)
	sort.Ints(result)

	return result[B-1]

}
