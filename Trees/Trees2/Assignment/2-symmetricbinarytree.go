package Assignment

type treeNode struct {
	left *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode{
	var node *treeNode = new(treeNode)
	node.value=val
	node.left=nil
	node.right=nil
	return node
}


func isSymmetric(A *treeNode )  (int) {

	if isMirror(A, A) {
		return 1
	}
	return 0
}


func isMirror(root1 *treeNode, root2 *treeNode) bool {

	if root1 == nil && root2 == nil {
		return true
	}


	if root1 != nil && root2 != nil {
		return root1.value == root2.value && isMirror(root1.left, root2.right) && isMirror(root1.right, root2.left)
	}

	return false

}
