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

func postorderTraversal(A *TreeNode) []int {

	var result []int
	var postorderTraversal func(*TreeNode)

	postorderTraversal = func(A *TreeNode) {

		if A == nil {
			return
		}

		postorderTraversal(A.Left)
		postorderTraversal(A.Right)
		result = append(result, A.Value)
	}

	postorderTraversal(A)
	return result

}
