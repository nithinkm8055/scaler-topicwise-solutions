package Homework

//Definition for binary tree
type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

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
