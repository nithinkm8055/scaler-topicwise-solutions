package Assignment

//type TreeNode struct {
//	Left  *TreeNode
//	Value int
//	Right *TreeNode
//}
//
//func TreeNode_new(val int) *TreeNode {
//	var node *TreeNode = new(TreeNode)
//	node.Value = val
//	node.Left = nil
//	node.Right = nil
//	return node
//}

func preorderTraversal(A *TreeNode) []int {

	var result []int

	var preorderTraversal func(*TreeNode)

	preorderTraversal = func(A *TreeNode) {

		if A == nil {
			return
		}
		result = append(result, A.Value)
		preorderTraversal(A.Left)
		preorderTraversal(A.Right)
	}

	preorderTraversal(A)

	return result
}
