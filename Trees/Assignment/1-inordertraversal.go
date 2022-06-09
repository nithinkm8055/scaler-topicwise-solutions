package Assignment

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

func TreeNode_new(val int) *TreeNode {
	var node *TreeNode = new(TreeNode)
	node.Value = val
	node.Left = nil
	node.Right = nil
	return node
}

func InorderTraversal(A *TreeNode) []int {

	var result []int
	var inorderTraversal func(*TreeNode)

	inorderTraversal = func(A *TreeNode) {
		if A == nil {
			return
		}

		inorderTraversal(A.Left)
		result = append(result, A.Value)
		inorderTraversal(A.Right)
	}

	inorderTraversal(A)

	return result

}
