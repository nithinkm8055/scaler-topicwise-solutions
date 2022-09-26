package Homework

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func t2Sum(A *treeNode, B int) int {

	complementMap := make(map[int]bool)

	var preorder func(*treeNode, int) int
	var ans int

	preorder = func(A *treeNode, B int) int {

		if A == nil {
			return 0
		}

		if _, ok := complementMap[A.value]; ok {
			ans = 1
			return ans
		} else {
			complementMap[B-A.value] = true
		}

		preorder(A.left, B)
		preorder(A.right, B)

		return ans
	}

	return preorder(A, B)
}
